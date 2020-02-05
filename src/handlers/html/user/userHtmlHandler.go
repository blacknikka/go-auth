package html

import (
	"log"
	"net/http"
	"text/template"

	"github.com/blacknikka/go-auth/usecases"
)

// UserHTMLHandler UserのHtmlハンドラ
type UserHTMLHandler interface {
	ShowAllUsers(http.ResponseWriter, *http.Request)
}

type userHTMLHandler struct {
	userUseCase usecases.UserUseCase
}

// NewUserHTMLHandler UserのHTMLハンドラを返す
func NewUserHTMLHandler(uu usecases.UserUseCase) UserHTMLHandler {
	return &userHTMLHandler{
		userUseCase: uu,
	}
}

func (uh userHTMLHandler) ShowAllUsers(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	if err != nil {
		log.Fatal(err)
	}

	if err := t.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>"); err != nil {
		log.Fatal(err)
	}
}
