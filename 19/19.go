package main

import (
	"fmt"
)

// Функция для переворачивания строки
func reverseString(s string) string {
	runes := []rune(s) // Преобразуем строку в срез рун
	n := len(runes)
	// Меняем местами символы (первый с последним, двигаясь к середине)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-i-1] = runes[n-i-1], runes[i]
	}
	return string(runes) // Преобразуем срез рун обратно в строку
}

func main() {
	input := "абырвалг"
	output := reverseString(input)
	fmt.Println(output) // Вывод: главрыба
}
