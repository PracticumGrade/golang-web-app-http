// авторское решение из модуля 2 урока 1 темы 1.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите JSON-данные:")

	// Создаем сканер для считывания ввода пользователя с консоли
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputJSON := scanner.Bytes()

	// Проверяем валидность JSON
	if json.Valid(inputJSON) {
		fmt.Println("Данные JSON являются валидными.")

		return
	}

	fmt.Println("Данные JSON недействительны.")
}
