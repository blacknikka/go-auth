package file

import (
	"net/http"
	"net/http/httptest"
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
	t.Run("UploadFile正常系", func(t *testing.T) {
		hasBeenCalled := false

		// DI (UserUseCase)
		spy := &fakeFileUseCase{
			FakeCreate: func(files.File) (*files.File, error) {
				mockedFileData := files.File{
					ID:   1,
					Name: "filename.jpg",
					Data: []byte{0x01, 0x02},
				}

				hasBeenCalled = true

				return &mockedFileData, nil
			},
		}

		fileHandler := NewFileHTMLHandler(spy)

		req, err := http.NewRequest("POST", "/file/upload", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(fileHandler.UploadFile)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if hasBeenCalled != true {
			t.Error("The function should be called")
		}
	})
}
