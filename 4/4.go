package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Функция для вывода на экран результата
func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Воркер №%d получил работу №%d\n", id, job)
	}
}

func main() {
	// С помощью флага в терминале можно назначить кол-во воркеров (5 по умолчанию)
	numWorkers := flag.Int("numworkers", 5, "Количество используемых воркеров")
	flag.Parse() // Чтение флага (если он есть)
	jobs := make(chan int)
	var wg sync.WaitGroup
	// Добавляем контекст для закрытия программы
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Запускаем воркеров
	for i := 1; i <= *numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Генерация данных для канала
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(jobs) // Завершение работы через контекст
				return
			default:
				// Раз в секунду в канал посупает случайное число (от 0 до 100)
				jobs <- rand.Intn(100)
				time.Sleep(time.Second)
			}
		}
	}()

	// Обработка сигнала завершения
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	<-sigCh
	fmt.Println("\nЗавершение работы программы...")
	cancel()
	wg.Wait()
}
