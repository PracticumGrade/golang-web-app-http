package main

import (
	"basic-rest-with-tests/internal/rest"
	"net/http"
)

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			rest.GetTasks(w, r)
		case http.MethodPost:
			rest.CreateTask(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			rest.UpdateTask(w, r)
		case http.MethodDelete:
			rest.DeleteTask(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
