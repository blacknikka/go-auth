package file

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"strconv"
	"unsafe"

	"github.com/blacknikka/go-auth/domain/models/files"
	fileUsecase "github.com/blacknikka/go-auth/usecases/file"
)

// FileHTMLHandler FileのHtmlハンドラ
type FileHTMLHandler interface {
	FindFileForm(http.ResponseWriter, *http.Request)
	FindFileRequest(http.ResponseWriter, *http.Request)
	UploadFileForm(http.ResponseWriter, *http.Request)
	UploadFileRequest(http.ResponseWriter, *http.Request)
}

type fileHTMLHandler struct {
	fileUsecase fileUsecase.FileUseCase
}

// NewFileHTMLHandler FileのHTMLハンドラを返す
func NewFileHTMLHandler(fu fileUsecase.FileUseCase) FileHTMLHandler {
	return &fileHTMLHandler{
		fileUsecase: fu,
	}
}

func (fh fileHTMLHandler) FindFileForm(w http.ResponseWriter, r *http.Request) {
	findTemplate := "./handlers/html/templates/findAPicture.html.tpl"

	t, err := template.ParseFiles(findTemplate)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		log.Fatal(err)
	}

	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func (fh fileHTMLHandler) FindFileRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.ParseForm()

	resultTemplate := "./handlers/html/templates/findAPictureResult.html.tpl"

	t, err := template.ParseFiles(resultTemplate)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		log.Fatal(err)
	}

	id := r.Form["file_id"][0]

	// call create function.
	intID, _ := strconv.Atoi(id)
	file, err := fh.fileUsecase.FindByID(intID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "file not found: %v", err)
		return
	}

	data := struct {
		Name string
		IMG  string
	}{
		Name: file.Name,
		IMG:  base64.StdEncoding.EncodeToString(file.Data),
	}

	if err := t.Execute(w, data); err != nil {
		log.Fatal(err)
	}
}

func (fh fileHTMLHandler) UploadFileForm(w http.ResponseWriter, r *http.Request) {
	uploadTemplate := "./handlers/html/templates/uploadAPicture.html.tpl"

	t, err := template.ParseFiles(uploadTemplate)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		log.Fatal(err)
	}

	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}

func (fh fileHTMLHandler) UploadFileRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.ParseForm()

	file, fileHeader, err := r.FormFile("up_data")
	defer file.Close()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "file is needed")
		return
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "save file failed.")
		return
	}

	fmt.Fprintf(w, "file: %v", fileHeader)

	// call create function.
	fh.fileUsecase.Create(files.File{Name: fileHeader.Filename, Data: *(*[]byte)(unsafe.Pointer(buf))})
}
