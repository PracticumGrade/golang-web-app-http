//file:testwebapp/internal/handlers/simpleprinter/handlers_test.go

package simpleprinter

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_mainPage(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	mainPage(w, r)

	result := w.Result()

	// проверяем код состояния
	if result.StatusCode != http.StatusOK {
		t.Errorf("Status code expected 200 OK, got %s", http.StatusText(result.StatusCode))
	}

	buf, err := io.ReadAll(result.Body)
	if err != nil {
		t.Errorf("io.ReadAll unexpected error %s", err.Error())
	}

	if result.Body != nil {
		result.Body.Close() // тело запроса является io.ReadCloser, его следует закрывать, чтобы освободить ресурсы
	}

	// проверяем тело запроса
	if string(buf) != "Hello!\n" {
		t.Errorf("Response body expected \"Hello!\n\", got \"%s\"", string(buf))
	}
}

func Test_apiPage(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api", nil)

	apiPage(w, r)

	result := w.Result()

	// проверяем код состояния
	if result.StatusCode != http.StatusOK {
		t.Errorf("Status code expected 200 OK, got %s", http.StatusText(result.StatusCode))
	}

	buf, err := io.ReadAll(result.Body)
	if err != nil {
		t.Errorf("io.ReadAll unexpected error %s", err.Error())
	}

	if result.Body != nil {
		result.Body.Close() // тело запроса является io.ReadCloser, его следует закрывать, чтобы освободить ресурсы
	}

	// проверяем тело запроса
	if string(buf) != "This is \"/api\" page.\n" {
		t.Errorf("Response body expected \"This is \"/api\" page.\n\", got \"%s\"", string(buf))
	}
}
