package file

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/blacknikka/go-auth/domain/models/files"
	fileUsecase "github.com/blacknikka/go-auth/usecases/file"
)

// for mock
type fakeFileUseCase struct {
	fileUsecase.FileUseCase
	FakeCreate   func(files.File) (*files.File, error)
	FakeFindByID func(int) (*files.File, error)
}

func (s *fakeFileUseCase) Create(file files.File) (*files.File, error) {
	return s.FakeCreate(file)
}

func (s *fakeFileUseCase) FindByID(id int) (*files.File, error) {
	return s.FakeFindByID(id)
}

func TestFileHTMLHandler(t *testing.T) {
	t.Run("UploadFile異常系_POSTではない", func(t *testing.T) {
		spy := &fakeFileUseCase{}
		fileHandler := NewFileHTMLHandler(spy)
		req, err := http.NewRequest("GET", "/file/upload_request", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(fileHandler.UploadFileRequest)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusBadRequest {
			t.Errorf("handler should return bad request: got %v want %v",
				status, http.StatusBadRequest)
		}
	})

	t.Run("UploadFile正常系", func(t *testing.T) {
		var b bytes.Buffer

		// multipartのリクエスト用オブジェクト
		w := multipart.NewWriter(&b)

		// ハンドラの準備（テスト用リクエスト）
		spy := &fakeFileUseCase{
			FakeCreate: func(files.File) (*files.File, error) {
				return nil, nil
			},
		}
		fileHandler := NewFileHTMLHandler(spy)
		ts := httptest.NewServer(http.HandlerFunc(fileHandler.UploadFileRequest))
		defer ts.Close()
		client := ts.Client()

		var fw io.Writer
		f, err := os.Open("main.go")
		if err != nil {
			fmt.Errorf("err should be nil: %v", err)
		}

		w.CreateFormFile("up_data", "main.go")
		written, err := io.Copy(fw, f)
		if written == 0 {
			fmt.Errorf("written should be a positive value: %v", err)
		}

		w.Close()

		req, err := http.NewRequest("POST", ts.URL, &b)
		if err != nil {
			fmt.Errorf("err should be nil: %v", err)
		}
		// Don't forget to set the content type, this will contain the boundary.
		req.Header.Set("Content-Type", w.FormDataContentType())

		// Submit the request
		res, err := client.Do(req)
		if err != nil {
			fmt.Errorf("err should be nil: %v", err)
		}

		// Check the response
		if res.StatusCode != http.StatusOK {
			fmt.Errorf("bad status: %s", res.Status)
		}
	})
}
