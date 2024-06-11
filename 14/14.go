package main

import (
	"fmt"
)

// Функция для определения типа поступающей переменной
func determineType(i interface{}) {
	switch v := i.(type) { // Обработка возможных типов через кейсы
	case int:
		fmt.Printf("Тип int, значение %d\n", v)
	case string:
		fmt.Printf("Тип string, значение %s\n", v)
	case bool:
		fmt.Printf("Тип bool, значение %t\n", v)
	case chan int:
		fmt.Println("Тип chan int")
	default:
		fmt.Println("Неизвестный тип")
	}
}

func main() {
	determineType(42)
	determineType("hello")
	determineType(true)
	ch := make(chan int)
	determineType(ch)
}
