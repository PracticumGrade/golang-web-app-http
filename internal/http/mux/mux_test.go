package mux_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PracticumGrade/web-app-http/internal/http/mux"
)

func Test_New(t *testing.T) {
	m := mux.New("", "")

	endpoints := []string{
		"/printer/hello",
		"/printer/api",
		"/printer/version",
		"/printer/randomnumber",
		"/restapi/users",
		"/restapi/users/filtered",
	}

	for _, endpoint := range endpoints {
		t.Run(endpoint, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, endpoint, nil)

			m.ServeHTTP(w, r)

			if w.Result().StatusCode == http.StatusNotFound {
				t.Errorf("endpoint %s not found in mux", endpoint)
			}
		})
	}
}
