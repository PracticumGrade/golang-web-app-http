package versionprinter_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PracticumGrade/web-app-http/internal/module1/http/handlers/versionprinter"
)

func Test_NewVersionPrinter(t *testing.T) {
	vp := versionprinter.NewVersionPrinter("", "")
	if vp == nil {
		t.Fatalf("versionprinter.NewVersionPrinter returned unexpected nil")
	}
}

func Test_VersionPrint_ServeHTTP(t *testing.T) {
	const (
		version   = "rolling"
		buildTime = "yesterday"
	)

	vp := versionprinter.NewVersionPrinter(version, buildTime)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	vp.ServeHTTP(w, r)

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
	if string(buf) != version+"_"+buildTime+"\n" {
		t.Errorf(`Response body expected "`+version+`_`+buildTime+`\n", got \"%s\"`, string(buf))
	}
}
