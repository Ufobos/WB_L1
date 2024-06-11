package main

import (
	"fmt"
	"strings"
)

// Функция проверки уникальности символов в строке
func isUnique(str string) bool {
	str = strings.ToLower(str) // Приведение строки к нижнему регистру
	charSet := make(map[rune]bool)

	// Циклом проходим строку, добавляя элементы в карту
	for _, char := range str {
		if _, exists := charSet[char]; exists { // Если в карте уже есть такой элемент,
			return false // возвращаем false
		}
		charSet[char] = true
	}
	return true
}

func main() {
	tests := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, test := range tests {
		fmt.Printf("Строка %s уникальна: %v\n", test, isUnique(test))
	}
}
