package main

import "fmt"

// Определяем структуру Human с произвольными полями и методами
type Human struct {
	Name string
	Age  int
}

func (h *Human) Speak() {
	fmt.Printf("Привет, меня зовут %s и мне %d лет.\n", h.Name, h.Age)
}

// Определяем структуру Action, встраивая структуру Human
type Action struct {
	Human
	Activity string
}

func main() {
	// Создаем экземпляр структуры Action
	action := Action{
		Human: Human{
			Name: "Даниил",
			Age:  25,
		},
		Activity: "бегает",
	}

	// Вызываем метод Speak, который встроен от Human
	action.Speak()

	// Используем собственное поле Activity
	fmt.Printf("%s %s.\n", action.Name, action.Activity)
}
