package main

import (
	"fmt"
)

// Функция для разделения массива на две части
func partition(arr []int, low, high int) int {
	pivot := arr[high] // выбираем последний элемент в качестве опорного
	i := low - 1       // индекс меньшего элемента

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i] // обмен элементов
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Рекурсивная функция быстрой сортировки
func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high) // индекс опорного элемента

		// рекурсивно сортируем элементы до и после опорного элемента
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func main() {
	arr := []int{10, 6, 2, 8, 9, 1, 5}
	n := len(arr) - 1
	quickSort(arr, 0, n)
	fmt.Println("Отсортированный массив:", arr)
}

// // Функция быстрой сортировки с использованием sort.Slice
// func quickSort(arr []int) {
//     // В sort.Slice передаётся срез и функция, определяющая порядок сортировки
//     sort.Slice(arr, func(i, j int) bool {
//         return arr[i] < arr[j]
//     })
// }
