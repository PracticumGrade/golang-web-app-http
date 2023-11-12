package filteredrest_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PracticumGrade/web-app-http/internal/modules2and3/http/handlers/filteredrest"

	"github.com/PracticumGrade/web-app-http/internal/modules2and3/storage"
)

func Test_UsersManagement_ServeHTTP_getUsers(t *testing.T) {
	// Создайте фейковое хранилище.
	fakeStorage := storage.New(".")
	fakeStorage.Store(storage.UserData{
		UserID:    1,
		UserName:  "User1",
		UserEmail: "user1@example.com",
	})
	fakeStorage.Store(storage.UserData{
		UserID:    2,
		UserName:  "User2",
		UserEmail: "user2@example.com",
	})
	fakeStorage.Store(storage.UserData{
		UserID:    3,
		UserName:  "User3",
		UserEmail: "user3@example.com",
	})
	fakeStorage.Store(storage.UserData{
		UserID:    4,
		UserName:  "User4",
		UserEmail: "user4@example.com",
	})

	usersHandler := filteredrest.NewFilteredRESTHandler(fakeStorage)

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
			request:    "/books?page=1&limit=2",
			statusCode: http.StatusOK,
			checkResponse: func(responseBody io.Reader) error {
				var (
					usersFromResponse []filteredrest.User
				)

				// Парсим JSON
				err := json.NewDecoder(responseBody).Decode(&usersFromResponse)
				if err != nil {
					return err
				}

				// Проверяем длину списка книг
				if len(usersFromResponse) != 2 {
					return fmt.Errorf("handler returned unexpected number of books: got %v want 2",
						len(usersFromResponse))
				}

				return nil
			},
		},
		{
			name:       "FilterBySpecificName",
			request:    "/books?name=User1",
			statusCode: http.StatusOK,
			checkResponse: func(responseBody io.Reader) error {
				var (
					usersFromResponse []filteredrest.User
				)

				// Парсим JSON
				err := json.NewDecoder(responseBody).Decode(&usersFromResponse)
				if err != nil {
					return err
				}

				// Проверяем длину списка книг
				if len(usersFromResponse) != 1 {
					return fmt.Errorf("handler returned unexpected number of books: got %v want 1",
						len(usersFromResponse))
				}

				return err
			},
		},
		{
			name:       "FilterByPartialName",
			request:    "/books?name=User",
			statusCode: http.StatusOK,
			checkResponse: func(responseBody io.Reader) error {
				var (
					usersFromResponse []filteredrest.User
				)

				// Парсим JSON
				err := json.NewDecoder(responseBody).Decode(&usersFromResponse)
				if err != nil {
					return err
				}

				// Проверяем длину списка книг
				if len(usersFromResponse) != 4 {
					return fmt.Errorf("handler returned unexpected number of books: got %v want 4",
						len(usersFromResponse))
				}

				return err
			},
		},
		{
			name:       "FilterBySpecificEmail",
			request:    "/books?email=user1@example.com",
			statusCode: http.StatusOK,
			checkResponse: func(responseBody io.Reader) error {
				var (
					usersFromResponse []filteredrest.User
				)

				// Парсим JSON
				err := json.NewDecoder(responseBody).Decode(&usersFromResponse)
				if err != nil {
					return err
				}

				// Проверяем длину списка книг
				if len(usersFromResponse) != 1 {
					return fmt.Errorf("handler returned unexpected number of books: got %v want 1",
						len(usersFromResponse))
				}

				return err
			},
		},
		{
			name:       "FilterByPartialEmail",
			request:    "/books?email=example.com",
			statusCode: http.StatusOK,
			checkResponse: func(responseBody io.Reader) error {
				var (
					usersFromResponse []filteredrest.User
				)

				// Парсим JSON
				err := json.NewDecoder(responseBody).Decode(&usersFromResponse)
				if err != nil {
					return err
				}

				// Проверяем длину списка книг
				if len(usersFromResponse) != 4 {
					return fmt.Errorf("handler returned unexpected number of books: got %v want 4",
						len(usersFromResponse))
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

			// Вызываем ServeHTTP с фейковым запросом и записываем ответ в ResponseRecorder.
			usersHandler.ServeHTTP(rr, req)

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
