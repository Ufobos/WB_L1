package main

import (
	"fmt"
)

// Функция для записи чисел в канал
func sendNumbers(numbers []int, inChan chan<- int) {
	for _, num := range numbers {
		inChan <- num
	}
	close(inChan) // Закрываем канал после отправки всех чисел
}

// Функция для чтения чисел из канала,
// умножения их на 2 и записи результатов во второй канал
func multiplyByTwo(inChan <-chan int, outChan chan<- int) {
	for num := range inChan {
		outChan <- num * 2
	}
	close(outChan) // Закрываем канал после обработки всех чисел
}

func main() {
	// Массив чисел
	numbers := []int{1, 2, 3, 4, 5}

	// Создаем два канала
	inChan := make(chan int)
	outChan := make(chan int)

	// Запускаем горутину для записи чисел в первый канал
	go sendNumbers(numbers, inChan)

	// Запускаем горутину для обработки чисел и записи результатов во второй канал
	go multiplyByTwo(inChan, outChan)

	// Читаем результаты из второго канала и выводим их на экран
	for result := range outChan {
		fmt.Println(result)
	}
}
