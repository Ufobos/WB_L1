package main

import "fmt"

// Функция для создания множества из последовательности строк
func createSet(strings []string) map[string]struct{} {
	// Создаем карту для хранения множества
	set := make(map[string]struct{})
	for _, str := range strings {
		// Если такой ключ уже есть, он перезаписывается
		set[str] = struct{}{} // Добавляем строки в карту
	}
	return set
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	set := createSet(strings)
	fmt.Println(set) // Выводим множество: map[cat:{} dog:{} tree:{}]
}
