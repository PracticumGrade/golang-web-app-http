// авторское решение из модуля 2 урока 3 темы 1.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Email string `json:"Email"`
}

func main() {
	// Чтение JSON-файла
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	// Создание среза пользователей для десериализации
	var users []User

	// Десериализация JSON-данных в структуры User
	if err := json.Unmarshal(data, &users); err != nil {
		fmt.Println("Ошибка десериализации данных:", err)
		return
	}

	// Вывод информации о пользователях
	for i, user := range users {
		fmt.Printf("User %d:\n", i+1)
		fmt.Printf("Name: %s\n", user.Name)
		fmt.Printf("Age: %d\n", user.Age)
		fmt.Printf("Email: %s\n", user.Email)
		fmt.Println()
	}
}
