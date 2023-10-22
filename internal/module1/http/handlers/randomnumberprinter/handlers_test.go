package randomnumberprinter_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PracticumGrade/web-app-http/internal/module1/http/handlers/randomnumberprinter"
)

func Test_NewVersionPrinter(t *testing.T) {
	rnp := randomnumberprinter.NewRandomNumberPrinter(nil)
	if rnp == nil {
		t.Fatalf("randomnumberprinter.NewRandomNumberPrinter returned unexpected nil")
	}
}

type randNumpProvider struct{}

func (randNumpProvider) GetRandomInt() int {
	return 1234
}

func Test_VersionPrint_ServeHTTP(t *testing.T) {
	rnp := randomnumberprinter.NewRandomNumberPrinter(randNumpProvider{})

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	rnp.ServeHTTP(w, r)

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
	if string(buf) != "1234\n" {
		t.Errorf(`Response body expected "1234\n", got \"%s\"`, string(buf))
	}
}
