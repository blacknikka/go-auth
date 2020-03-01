package file

import (
	"fmt"
	"net/http"

	fileUsecase "github.com/blacknikka/go-auth/usecases/file"
)

// FileHTMLHandler FileのHtmlハンドラ
type FileHTMLHandler interface {
	UploadFile(http.ResponseWriter, *http.Request)
}

type fileHTMLHandler struct {
	fileUseCase fileUsecase.FileUseCase
}

// NewFileHTMLHandler FileのHTMLハンドラを返す
func NewFileHTMLHandler(fu fileUsecase.FileUseCase) FileHTMLHandler {
	return &fileHTMLHandler{
		fileUseCase: fu,
	}
}

func (fh fileHTMLHandler) UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
