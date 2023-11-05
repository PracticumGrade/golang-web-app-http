package rest_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PracticumGrade/web-app-http/internal/modules2and3/http/handlers/rest"
	"github.com/PracticumGrade/web-app-http/internal/modules2and3/storage"
)

func Test_UsersManagement_ServeHTTP_getUsers(t *testing.T) {
	// Создайте фейковое хранилище и добавьте в него тестовых пользователей.
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

	// Создайте HTTP-запрос GET /users.
	req := httptest.NewRequest("GET", "/restapi/users", nil)
	w := httptest.NewRecorder()

	// Создайте обработчик REST и выполните запрос.
	handler := rest.NewRESTHandler(fakeStorage)
	handler.ServeHTTP(w, req)

	// Проверьте, что получен правильный статус ответа и JSON-список пользователей.
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	var users []storage.UserData
	err := json.NewDecoder(w.Body).Decode(&users)
	if err != nil {
		t.Errorf("Error decoding response: %v", err)
	}

	// Проверьте, что список пользователей не пустой и содержит ожидаемых пользователей.
	if len(users) != 2 {
		t.Errorf("Expected 2 users, got %d", len(users))
	}
}

func Test_UsersManagement_ServeHTTP_createUser(t *testing.T) {
	// Создайте фейковое хранилище.
	fakeStorage := storage.New(".")

	// Создайте нового пользователя для создания.
	newUser := storage.UserData{
		UserName:  "NewUser",
		UserEmail: "newuser@example.com",
	}

	// Преобразуйте пользователя в JSON.
	userJSON, err := json.Marshal(newUser)
	if err != nil {
		t.Errorf("Error encoding user: %v", err)
	}

	// Создайте HTTP-запрос POST /users с телом запроса, содержащим пользователя в JSON.
	req := httptest.NewRequest("POST", "/restapi/users", bytes.NewReader(userJSON))
	w := httptest.NewRecorder()

	// Создайте обработчик REST и выполните запрос.
	handler := rest.NewRESTHandler(fakeStorage)
	handler.ServeHTTP(w, req)

	// Проверьте, что получен правильный статус ответа.
	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}

	// Проверьте, что созданный пользователь имеет поле UserID (идентификатор пользователя).
	var createdUser storage.UserData
	err = json.NewDecoder(w.Body).Decode(&createdUser)
	if err != nil {
		t.Errorf("Error decoding response: %v", err)
	}
	if createdUser.UserID == 0 {
		t.Errorf("Expected non-zero UserID, got 0")
	}

	// Проверьте, что пользователь действительно создан в хранилище.
	storedUser, err := fakeStorage.Recover(createdUser.UserID)
	if err != nil {
		t.Errorf("Error recovering user from storage: %v", err)
	}
	if storedUser.UserName != createdUser.UserName || storedUser.UserEmail != createdUser.UserEmail {
		t.Errorf("Stored user doesn't match created user")
	}
}

func Test_UsersManagement_ServeHTTP_updateUser(t *testing.T) {
	// Создайте фейковое хранилище и добавьте в него тестового пользователя.
	fakeStorage := storage.New(".")
	userID, _ := fakeStorage.Store(storage.UserData{
		UserName:  "UserToUpdate",
		UserEmail: "update@example.com",
	})

	// Подготовьте обновленные данные пользователя.
	updatedUser := storage.UserData{
		UserID:    userID,
		UserName:  "UpdatedUser",
		UserEmail: "updated@example.com",
	}

	// Преобразуйте обновленные данные пользователя в JSON.
	userJSON, err := json.Marshal(updatedUser)
	if err != nil {
		t.Errorf("Error encoding user: %v", err)
	}

	// Создайте HTTP-запрос PUT /users с телом запроса, содержащим обновленные данные пользователя в JSON.
	req := httptest.NewRequest("PUT", "/restapi/users", bytes.NewReader(userJSON))
	w := httptest.NewRecorder()

	// Создайте обработчик REST и выполните запрос.
	handler := rest.NewRESTHandler(fakeStorage)
	handler.ServeHTTP(w, req)

	// Проверьте, что получен правильный статус ответа.
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Проверьте, что обновленные данные пользователя соответствуют ожидаемым.
	var responseUser storage.UserData
	err = json.NewDecoder(w.Body).Decode(&responseUser)
	if err != nil {
		t.Errorf("Error decoding response: %v", err)
	}
	if responseUser != updatedUser {
		t.Errorf("Updated user data doesn't match the expected data")
	}

	// Проверьте, что пользователь в хранилище действительно обновлен.
	storedUser, err := fakeStorage.Recover(userID)
	if err != nil {
		t.Errorf("Error recovering user from storage: %v", err)
	}
	if storedUser != updatedUser {
		t.Errorf("Stored user doesn't match updated user data")
	}
}

func Test_UsersManagement_ServeHTTP_deleteUser(t *testing.T) {
	// Создайте фейковое хранилище и добавьте в него тестового пользователя.
	fakeStorage := storage.New(".")
	userID, _ := fakeStorage.Store(storage.UserData{
		UserName:  "UserToDelete",
		UserEmail: "delete@example.com",
	})

	// Преобразуйте ID пользователя в JSON.
	userJSON, err := json.Marshal(storage.UserData{UserID: userID})
	if err != nil {
		t.Errorf("Error encoding user: %v", err)
	}

	// Создайте HTTP-запрос DELETE /users/{id}, где {id} - идентификатор пользователя.
	req := httptest.NewRequest("DELETE", "/restapi/users", bytes.NewReader(userJSON))
	w := httptest.NewRecorder()

	// Создайте обработчик REST и выполните запрос.
	handler := rest.NewRESTHandler(fakeStorage)
	handler.ServeHTTP(w, req)

	// Проверьте, что получен правильный статус ответа.
	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	// Проверьте, что пользователь действительно удален из хранилища.
	_, err = fakeStorage.Recover(userID)
	if !errors.Is(err, storage.ErrNotFound) {
		t.Errorf("Expected user to be deleted, but it still exists in storage")
	}
}
