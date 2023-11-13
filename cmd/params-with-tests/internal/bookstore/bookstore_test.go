//file:<project>/internal/bookstore/bookstore_test.go

package bookstore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_BooksHandler(t *testing.T) {
	// Создаем таблицу тестирования
	testingTable := []struct {
		// Наименование теста
		name string
		// Запрос
		request string
		// Ожидаемый статус ответа
		statusCode int
		// Лямбда, проверяющая тело ответа. Возвращает ошибку, если проверка не пройдена.
		checkResponse func(responseBody io.Reader) error
	}{
		{
			name:       "Pagination",
			request:    "/books?page=2&limit=10",
			statusCode: http.StatusOK,
			checkResponse: func(responseBody io.Reader) error {
				var (
					booksFromResponse []Book
				)

				// Парсим JSON
				err := json.NewDecoder(responseBody).Decode(&booksFromResponse)
				if err != nil {
					return err
				}

				// Проверяем длину списка книг
				if len(booksFromResponse) != 10 {
					return fmt.Errorf("handler returned unexpected number of books: got %v want 10",
						len(booksFromResponse))
				}

				return nil
			},
		},
		{
			name:       "FilterByAuthor",
			request:    "/books?author=Author%202", // %20 - пробел
			statusCode: http.StatusOK,
			checkResponse: func(responseBody io.Reader) error {
				var (
					booksFromResponse []Book
				)

				// Парсим JSON
				err := json.NewDecoder(responseBody).Decode(&booksFromResponse)
				if err != nil {
					return err
				}

				// Проверяем длину списка книг
				if len(booksFromResponse) != 17 {
					return fmt.Errorf("handler returned unexpected number of books: got %v want 10",
						len(booksFromResponse))
				}

				for _, v := range booksFromResponse {
					if v.Author != "Author 2" {
						err = errors.Join(err,
							fmt.Errorf("book with id %d has unexpected author name: got %s, want Author 2",
								v.ID, v.Author))
					}
				}

				return err
			},
		},
	}

	for _, member := range testingTable {
		t.Run(member.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", member.request, nil)

			// Создаем ResponseRecorder для записи ответа.
			rr := httptest.NewRecorder()

			// Создаем хэндлер с тестируемой функцией.
			handler := http.HandlerFunc(BooksHandler)

			// Вызываем ServeHTTP с фейковым запросом и записываем ответ в ResponseRecorder.
			handler.ServeHTTP(rr, req)

			// Проверяем код статуса.
			if status := rr.Code; status != member.statusCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, member.statusCode)
			}

			if err := member.checkResponse(rr.Body); err != nil {
				t.Errorf("response body check function returned unexpected error: %v", err)
			}
		})
	}
}
