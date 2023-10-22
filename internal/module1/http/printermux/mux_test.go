package printermux_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PracticumGrade/web-app-http/internal/module1/http/printermux"
)

func Test_New(t *testing.T) {
	mux := printermux.New("", "")

	endpoints := []string{
		"/hello",
		"/api",
		"/version",
		"/randomnumber",
	}

	for _, endpoint := range endpoints {
		t.Run(endpoint, func(t *testing.T) {
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, endpoint, nil)

			mux.ServeHTTP(w, r)

			if w.Result().StatusCode == http.StatusNotFound {
				t.Errorf("endpoint %s not found in mux", endpoint)
			}
		})
	}
}
