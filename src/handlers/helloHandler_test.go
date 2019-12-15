package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloController(t *testing.T) {
	t.Run("GET HelloServer", func(t *testing.T) {
		// https://blog.questionable.services/article/testing-http-handlers-go/
		// 参考に実装.

		req, err := http.NewRequest("GET", "/hello", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HelloServer)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := `hello, world!`
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})

	t.Run("POST HelloJSONHandle", func(t *testing.T) {
		jsonObject := JSONRequest{Name: "my-name"}
		jsonString, err := json.Marshal(jsonObject)
		if err != nil {
			t.Fatal(err)
		}

		req, err := http.NewRequest(
			"POST",
			"/hello",
			bytes.NewBuffer(jsonString))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(HelloJSONHandle)
		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected, err := json.Marshal(JSONResponse{Message: "hello my-name"})
		if err != nil {
			t.Fatal(err)
		}

		if rr.Body.String() != string(expected) {
			t.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	})
}
