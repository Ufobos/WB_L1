package main

import (
	"fmt"
	"strings"
)

// Функция для переворачивания слов в строке
func reverseWords(s string) string {
	words := strings.Fields(s) // Разбиваем строку на слова
	n := len(words)
	// Меняем местами слова
	for i := 0; i < n/2; i++ {
		words[i], words[n-i-1] = words[n-i-1], words[i]
	}
	return strings.Join(words, " ") // Собираем строку обратно
}

func main() {
	input := "snow dog sun"
	output := reverseWords(input)
	fmt.Println(output)
}
