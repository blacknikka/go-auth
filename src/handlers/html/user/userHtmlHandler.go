package user

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	userUsecase "github.com/blacknikka/go-auth/usecases/user"
)

// UserHTMLHandler UserのHtmlハンドラ
type UserHTMLHandler interface {
	ShowAllUsers(http.ResponseWriter, *http.Request)
}

type userHTMLHandler struct {
	userUseCase userUsecase.UserUseCase
}

// NewUserHTMLHandler UserのHTMLハンドラを返す
func NewUserHTMLHandler(uu userUsecase.UserUseCase) UserHTMLHandler {
	return &userHTMLHandler{
		userUseCase: uu,
	}
}

func (uh userHTMLHandler) ShowAllUsers(w http.ResponseWriter, r *http.Request) {
	fileForTemplate := "./handlers/html/templates/showAllUsers.html.tpl"

	t, err := template.ParseFiles(fileForTemplate)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		log.Fatal(err)
	}

	usersToShow, err := uh.userUseCase.GetAll()
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		log.Fatal(err)
	}

	if err := t.Execute(w, usersToShow); err != nil {
		log.Fatal(err)
	}
}
