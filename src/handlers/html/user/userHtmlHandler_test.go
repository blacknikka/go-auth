package html

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/blacknikka/go-auth/domain/models/users"
	userUsecase "github.com/blacknikka/go-auth/usecases/user"
)

// for mock
type fakeUserUseCase struct {
	userUsecase.UserUseCase
	FakeGetAll func() (*[]users.User, error)
}

func (s *fakeUserUseCase) GetAll() (*[]users.User, error) {
	return s.FakeGetAll()
}

func TestUserHtmlHandler(t *testing.T) {
	t.Run("ShowAllUsers正常系", func(t *testing.T) {
		targetPath, _ := filepath.Abs("../../../")
		os.Chdir(targetPath)
		fmt.Println(os.Getwd())

		hasBeenCalled := false

		// DI (UserUseCase)
		spy := &fakeUserUseCase{
			FakeGetAll: func() (*[]users.User, error) {
				mockedUsersData := []users.User{
					{
						ID:    1,
						Name:  "name",
						Email: "name@example.com",
					},
					{
						ID:    2,
						Name:  "ore",
						Email: "ore@example.com",
					},
				}

				hasBeenCalled = true

				return &mockedUsersData, nil
			},
		}

		userHandler := NewUserHTMLHandler(spy)

		req, err := http.NewRequest("GET", "/user/all", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.ShowAllUsers)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		if hasBeenCalled != true {
			t.Error("GetAll() should be called")
		}
	})
}
