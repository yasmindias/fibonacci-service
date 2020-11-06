package cmd

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		err, rr := makeRequest("GET", "/fib?n=10", nil)
		if err != nil {
			t.Fatal(err)
		}

		checkStatus(rr.Code, http.StatusOK, t)

		expected := "55"
		if rr.Body.String() != expected {
			t.Errorf("request returned unexpected body: got '%v' want '%v'",
				rr.Body.String(), expected)
		}
	})

	t.Run("bad request", func(t *testing.T) {
		err, rr := makeRequest("GET", "/fib?n=error", nil)
		if err != nil {
			t.Fatal(err)
		}

		checkStatus(rr.Code, http.StatusBadRequest, t)

		expected := "input must be an integer"
		if !strings.Contains(rr.Body.String(), expected) {
			t.Errorf("request returned unexpected body: got '%v' want '%v'",
				rr.Body.String(), expected)
		}
	})
}

func makeRequest(method, url string, body io.Reader) (error, *httptest.ResponseRecorder) {
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		return err, nil
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Get)
	handler.ServeHTTP(rr, req)

	return nil, rr
}

func checkStatus(received, expected int, t *testing.T) {
	if received != expected {
		t.Errorf("request returned wrong status code: got %v want %v",
			received, received)
	}
}
