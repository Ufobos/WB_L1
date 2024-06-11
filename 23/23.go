package main

import "fmt"

// Функция для удаления i-го элемента из слайса
// возвращает результат сложения двух частей без выбранного элемента
func removeElement(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный слайс:", slice)

	i := 2
	fmt.Println("Необходимо удалить элемент:", slice[i])
	slice = removeElement(slice, i)
	fmt.Println("Слайс после удаления элемента:", slice)
}
