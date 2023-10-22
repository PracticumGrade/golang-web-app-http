//file:testwebapp/internal/handlers/simpleprinter/mux_test.go

package simpleprinter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NewSimplePrinterServeMux(t *testing.T) {
	mux := NewSimplePrinterServeMux()

	endpoints := []string{
		"/",
		"/api",
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
