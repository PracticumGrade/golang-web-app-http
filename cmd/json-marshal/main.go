// авторское решение из модуля 2 урока 2 темы 1.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Book - структура для представления информации о книге
type Book struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Year   int    `json:"Year"`
}

func main() {
	var title, author string
	var year int

	// Создаем новый сканнер для стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)

	// Получаем информацию о книге от пользователя
	fmt.Print("Введите заголовок книги: ")
	scanner.Scan()
	title = scanner.Text()

	fmt.Print("Введите автора: ")
	scanner.Scan()
	author = scanner.Text()

	fmt.Print("Введите год публикации: ")
	scanner.Scan()
	_, err := fmt.Sscanf(scanner.Text(), "%d", &year)
	if err != nil {
		fmt.Println("Ошибка при считывании года:", err)

		return
	}

	// Создаем экземпляр структуры Book
	book := Book{
		Title:  title,
		Author: author,
		Year:   year,
	}

	// Преобразуем данные в JSON с отступами
	jsonData, err := json.MarshalIndent(book, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при маршалинге в JSON:", err)

		return
	}

	// Выводим отформатированный JSON в консоль
	fmt.Println(string(jsonData))
}
