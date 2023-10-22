//file:testwebapp/internal/server_test/server_test.go

package server_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testwebapp/internal/handlers/simpleprinter"
)

func TestHTTPServer(t *testing.T) {
	s := httptest.NewServer(simpleprinter.NewSimplePrinterServeMux())

	defer s.Close()

	// создаем таблицу тестирования с путем запроса и ожидаемым результатом в теле запроса
	type reqTable struct {
		path   string
		result string
	}

	table := []reqTable{
		{path: "/", result: "Hello!\n"},
		{path: "/api", result: "This is \"/api\" page.\n"},
	}

	for _, member := range table {
		t.Run(member.path, func(t *testing.T) {
			// создаем новый запрос
			req, err := http.NewRequest(http.MethodGet, s.URL+member.path, nil)
			if err != nil {
				t.Fatalf("NewRequest unexpected error %s", err.Error())
			}

			// осуществляем запрос
			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("DefaultClient.Do unexpected error %s", err.Error())
			}

			// проверяем код состояния
			if resp.StatusCode != http.StatusOK {
				t.Errorf("Status code expected 200 OK, got %s", http.StatusText(resp.StatusCode))
			}

			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("io.ReadAll unexpected error %s", err.Error())
			}

			if resp.Body != nil {
				resp.Body.Close() // тело запроса является io.ReadCloser, его следует закрывать, чтобы освободить ресурсы
			}

			// проверяем тело запроса
			if string(buf) != member.result {
				t.Errorf("Response body expected \"%s\n\", got \"%s\"", member.result, string(buf))
			}
		})
	}
}
