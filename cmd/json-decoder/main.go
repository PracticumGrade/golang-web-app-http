// авторское решение из модуля 2 урока 4 темы 1.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Year   int    `json:"Year"`
}

func main() {
	// Открываем файл "data.json" для чтения
	file, err := os.Open("books.json")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)

		return
	}

	defer file.Close()

	// Инициализируем декодер JSON
	decoder := json.NewDecoder(file)

	// Создаем срез для хранения книг
	var books []Book

	// Декодируем JSON и загружаем данные в срез
	if err := decoder.Decode(&books); err != nil {
		fmt.Println("Ошибка декодирования JSON:", err)

		return
	}

	// Выводим информацию о книгах
	fmt.Println("Список книг:")

	for _, book := range books {
		fmt.Printf("Заголовок: %s\n", book.Title)
		fmt.Printf("Автор: %s\n", book.Author)
		fmt.Printf("Год: %d\n", book.Year)
		fmt.Println("-----------")
	}
}
