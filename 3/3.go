package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	sumChan := make(chan int) // Канал для хранения квадратов чисел
	// Циклоим запускаем горутины, добавляющие квадраты чисел в канал
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			sumChan <- n * n
		}(num)
	}

	go func() {
		wg.Wait()      // Ждём добавления всех чисел в канал
		close(sumChan) // Закрываем канал
	}()

	sum := 0
	// Цикломи достаём квадраты чисел из канала и находим их сумму
	for sq := range sumChan {
		sum += sq
	}

	fmt.Printf("Сумма квадратов равна %d\n", sum)
}
