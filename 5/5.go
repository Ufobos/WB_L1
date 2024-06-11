package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// Определяем флаг для времени работы программы (по умолчанию это 10 секунд)
	duration := flag.Int("duration", 10, "Количество секунд работы программы")
	flag.Parse()

	values := make(chan int)
	// Горутина, отпрваляющая данные в канал (раз в секунду)
	go func() {
		for i := 0; ; i++ {
			values <- i
			time.Sleep(time.Second)
		}
	}()
	// Горутина для обработки записей из канала
	go func() {
		for v := range values {
			fmt.Printf("Получено значение: %d\n", v)
		}
	}()

	// Используем значение флага для ожидания
	time.Sleep(time.Duration(*duration) * time.Second)
	close(values)
	fmt.Println("Программа завершена.")
}
