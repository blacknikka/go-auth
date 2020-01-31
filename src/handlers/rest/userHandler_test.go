package rest

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blacknikka/go-auth/domain/models/users"
)

// for mock
type fakeUserUseCase struct {
	FakeGetAll func() (*[]users.User, error)
}

func (s *fakeUserUseCase) GetAll() (*[]users.User, error) {
	return s.FakeGetAll()
}

func TestUserController(t *testing.T) {
	t.Run("GET Index", func(t *testing.T) {
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

				return &mockedUsersData, nil
			},
		}

		userHandler := NewUserHandler(spy)

		req, err := http.NewRequest("GET", "/user", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Index)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		var responsedUsers Response
		if err = json.Unmarshal([]byte(rr.Body.String()), &responsedUsers); err != nil {
			t.Fatal(err)
		}

		if len(responsedUsers.Users) != 2 {
			t.Fatalf("Responsed data length should be 2, but %v", len(responsedUsers.Users))
		}

		userTests := []struct {
			name  string
			email string
		}{
			{name: "name", email: "name@example.com"},
			{name: "ore", email: "ore@example.com"},
		}

		for index, ut := range userTests {
			if responsedUsers.Users[index].Name != ut.name {
				t.Errorf("Name: got %v want %v", responsedUsers.Users[index].Name, "name")
			}

			if responsedUsers.Users[index].Email.Email != ut.email {
				t.Errorf("Email: got %v want %v", responsedUsers.Users[index].Email, "name@example.com")
			}
		}
	})

	t.Run("GET Index error pattern", func(t *testing.T) {
		spy := &fakeUserUseCase{
			FakeGetAll: func() (*[]users.User, error) {
				return nil, errors.New("connection failed.")
			},
		}

		// inject a mock object.
		userHandler := NewUserHandler(spy)

		req, err := http.NewRequest("GET", "/user", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(userHandler.Index)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusInternalServerError {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusInternalServerError)
		}

	})
}
