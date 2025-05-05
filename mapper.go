package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	p := Person{
		Name:  "Anna",
		Age:   25,
		Email: "anna-100mmmm@yandex.ru",
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Fatalf("Ошибка при сериализации: %v", err)
	}
	fmt.Println("Сериализованный JSON:", string(data))

	var p2 Person
	err = json.Unmarshal(data, &p2)
	if err != nil {
		log.Fatalf("Ошибка при десериализации: %v", err)
	}
	fmt.Println("Десериализованный объект:", p2)
}
