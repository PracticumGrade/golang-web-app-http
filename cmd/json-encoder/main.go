// авторское решение из модуля 2 урока 4 темы 1.

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
}

func main() {
	// Создаем несколько книг
	books := []Book{
		{"Book 1", "Author 1", 2022},
		{"Book 2", "Author 2", 2020},
		{"Book 3", "Author 1", 2019},
	}

	// Сериализуем массив в JSON
	jsonData, err := json.Marshal(books)
	if err != nil {
		fmt.Println("Ошибка сериализации JSON:", err)

		return
	}

	// Сохраняем JSON в файл
	file, err := os.Create("books.json")
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)

		return
	}

	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)

		return
	}

	fmt.Println("Данные успешно сохранены в файл 'books.json'.")

	// Теперь прочтем JSON из файла
	var loadedBooks []Book

	file, err = os.Open("books.json")
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)

		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&loadedBooks); err != nil {
		fmt.Println("Ошибка десериализации JSON:", err)

		return
	}

	fmt.Println("\nЗагруженные книги:")

	for _, book := range loadedBooks {
		fmt.Printf("Title: %s, Author: %s, Year: %d\n", book.Title, book.Author, book.Year)
	}
}
