package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTasks(t *testing.T) {
	// Создание тестового сервера
	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetTasks)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", status)
	}

	expectedContentType := "application/json"
	actualContentType := recorder.Header().Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected Content-Type: %s, but got %s", expectedContentType, actualContentType)
	}
}

func TestCreateTask(t *testing.T) {
	// Создание тестовой задачи
	newTask := Task{
		ID:     "1",
		Text:   "Sample Task",
		Status: "New",
	}

	taskJSON, err := json.Marshal(newTask)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTask)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusCreated {
		t.Errorf("Expected status 201, but got %d", status)
	}

	expectedContentType := "application/json"
	actualContentType := recorder.Header().Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected Content-Type: %s, but got %s", expectedContentType, actualContentType)
	}
}

func TestUpdateTask(t *testing.T) {
	// Создание тестовой задачи для обновления
	updateTask := Task{
		ID:     "1",
		Text:   "Updated Task",
		Status: "InProgress",
	}

	taskJSON, err := json.Marshal(updateTask)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(taskJSON))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateTask)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("Expected status 200, but got %d", status)
	}

	expectedContentType := "application/json"
	actualContentType := recorder.Header().Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected Content-Type: %s, but got %s", expectedContentType, actualContentType)
	}
}

func TestDeleteTask(t *testing.T) {
	// Создание тестовой задачи для удаления
	tasks = []Task{
		{
			ID:     "1",
			Text:   "Task 1",
			Status: "New",
		},
		{
			ID:     "2",
			Text:   "Task 2",
			Status: "InProgress",
		},
	}

	req, err := http.NewRequest("DELETE", "/tasks/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTask)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusNoContent {
		t.Errorf("Expected status 204, but got %d", status)
	}

	// Проверяем, что задача с ID "1" действительно удалена
	for _, task := range tasks {
		if task.ID == "1" {
			t.Error("Task with ID 1 was not deleted")
		}
	}
}

func TestDeleteTask_NotFound(t *testing.T) {
	// Создание тестовой задачи для удаления
	tasks = []Task{
		{
			ID:     "1",
			Text:   "Task 1",
			Status: "New",
		},
		{
			ID:     "2",
			Text:   "Task 2",
			Status: "InProgress",
		},
	}

	req, err := http.NewRequest("DELETE", "/tasks/3", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTask)

	handler.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusNotFound {
		t.Errorf("Expected status 404, but got %d", status)
	}
}
