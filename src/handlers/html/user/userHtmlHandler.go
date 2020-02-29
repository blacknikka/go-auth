package html

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/blacknikka/go-auth/domain/models/users"
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
	fileForTemplate := "./handlers/html/templates/showAllUsers.html.tpl"

	t, err := template.ParseFiles(fileForTemplate)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		log.Fatal(err)
	}

	usersToShow := []users.User{
		{ID: 1, Name: "Name1", Email: "email"},
		{ID: 2, Name: "Name2", Email: "email2"},
	}

	if err := t.Execute(w, usersToShow); err != nil {
		log.Fatal(err)
	}
}
