// авторское решение из модуля 3 урока 1 темы 1.

package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type responseT struct {
	Result     int    `json:"result"`
	Error      string `json:"error"`
	Parameters struct {
		Number1 int `json:"number1"`
		Number2 int `json:"number2"`
	} `json:"parameters"`
}

func calculateSum(a, b int) int {
	return a + b
}

func main() {
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		var response responseT

		// Извлекаем параметры из URL
		query := r.URL.Query()
		aStr := query.Get("number1")
		bStr := query.Get("number2")

		// Преобразуем параметры в числа
		number1, errA := strconv.Atoi(aStr)
		number2, errB := strconv.Atoi(bStr)

		response.Parameters.Number1 = number1
		response.Parameters.Number2 = number2

		// Проверяем ошибки
		if errA != nil || errB != nil {
			response.Error = "Invalid parameters"

			w.Header().Set("X-Content-Type-Options", "nosniff")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)

			return
		}

		// Вычисляем сумму
		result := calculateSum(number1, number2)
		response.Result = result

		// Кодируем ответ в JSON и отправляем
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	http.ListenAndServe(":8080", nil)
}
