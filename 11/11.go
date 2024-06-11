package main

import "fmt"

// Функция для нахождения пересечения двух множеств
func intersection(set1, set2 []int) []int {
	// Создаем карту для хранения элементов первого множества
	m := make(map[int]struct{})
	for _, v := range set1 {
		m[v] = struct{}{} // Добавляем элементы в карту
	}

	// Создаем срез для хранения результата пересечения
	var result []int
	for _, v := range set2 {
		if _, ok := m[v]; ok { // Проверяем, есть ли элемент из второго множества в карте
			result = append(result, v) // Если есть, добавляем его в результат
		}
	}

	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5, 5}
	set2 := []int{3, 4, 5, 6, 7}
	result := intersection(set1, set2)
	fmt.Println(result)
}
