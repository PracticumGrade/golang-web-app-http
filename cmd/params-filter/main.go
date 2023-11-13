// авторское решение из модуля 3 урока 1 темы 2.

package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	DueDate   time.Time `json:"dueDate"`
}

var tasks = []Task{
	{
		1,
		"Buy groceries",
		false,
		time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
	},
	{
		2,
		"Write code",
		true,
		time.Date(2023, time.November, 15, 0, 0, 0, 0, time.UTC),
	},
	{
		3,
		"Go for a run",
		false,
		time.Date(2023, time.December, 15, 0, 0, 0, 0, time.UTC),
	},
}

func filterTasksByCompleted(tasks []Task, completed bool) (filtered []Task) {
	for _, task := range tasks {
		if task.Completed == completed {
			filtered = append(filtered, task)
		}
	}

	return
}

func filterTasksByDueDate(tasks []Task, dueDate time.Time) (filtered []Task) {
	for _, task := range tasks {
		if dueDate.IsZero() || task.DueDate.Equal(dueDate) {
			filtered = append(filtered, task)
		}
	}

	return
}

func filterTasksByKeyword(tasks []Task, filter string) (filtered []Task) {
	for _, task := range tasks {
		if strings.Contains(task.Title, filter) {
			filtered = append(filtered, task)
		}
	}

	return
}

func main() {
	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		filteredTasks := tasks

		r.ParseForm()
		// Извлекаем параметр запроса 'completed'.
		completedParam := r.FormValue("completed")
		if completedParam != "" {
			// Преобразуем параметр в булево значение.
			completed, err := strconv.ParseBool(completedParam)
			if err != nil {
				http.Error(w, "Invalid parameter 'completed'", http.StatusBadRequest)

				return
			}

			// Инициируем фильтр задач на основе 'completed'.
			filteredTasks = filterTasksByCompleted(filteredTasks, completed)
		}

		// Извлекаем параметр запроса 'dueDate'.
		dueDateParam := r.FormValue("dueDate")
		if dueDateParam != "" {
			parsedDueDate, err := time.Parse("2006-01-02", dueDateParam)
			if err != nil {
				http.Error(w, "Invalid parameter 'dueDate'", http.StatusBadRequest)

				return
			}

			// Инициируем фильтр задач на основе 'dueDate'.
			filteredTasks = filterTasksByDueDate(filteredTasks, parsedDueDate)
		}

		filterParam := r.FormValue("filter")
		if filterParam != "" {
			filteredTasks = filterTasksByKeyword(filteredTasks, filterParam)
		}

		// Отправляем отфильтрованные задачи в формате JSON.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(filteredTasks)
	})

	http.ListenAndServe(":8080", nil)
}
