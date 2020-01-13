package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blacknikka/go-auth/domain/models/users"
)

// for mock
type spyUserUseCase struct {
}

func (s *spyUserUseCase) GetAll(context.Context) ([]*users.User, error) {
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

	return []*users.User{&mockedUsersData[0], &mockedUsersData[1]}, nil
}

func TestUserController(t *testing.T) {
	t.Run("GET Index", func(t *testing.T) {

		req, err := http.NewRequest("GET", "/user", nil)
		if err != nil {
			t.Fatal(err)
		}

		// DI (UserUseCase)
		spy := new(spyUserUseCase)
		userHandler := NewUserHandler(spy)

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
}
