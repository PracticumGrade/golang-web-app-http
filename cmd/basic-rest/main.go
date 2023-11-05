// авторское решение из модуля 2 урока 1 темы 2.

package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ID     string `json:"id"`
	Text   string `json:"text"`
	Status string `json:"status"`
}

var tasks []Task

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	tasks = append(tasks, task)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Path[len("/tasks/"):]

	var updatedTask Task
	_ = json.NewDecoder(r.Body).Decode(&updatedTask)

	for i, task := range tasks {
		if task.ID == taskID {
			tasks[i] = updatedTask
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Path[len("/tasks/"):]

	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			GetTasks(w, r)
		case http.MethodPost:
			CreateTask(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			UpdateTask(w, r)
		case http.MethodDelete:
			DeleteTask(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
