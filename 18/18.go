package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

// Метод для инкрементации счетчика
func (c *Counter) Increment() {
	c.mu.Lock() // Блокируем доступ к счетчику
	c.value++   // Инкрементируем значение
	fmt.Println("Текущее значение счётчика: ", c.value)
	c.mu.Unlock() // Разблокируем доступ к счетчику
}

// Метод для получения текущего значения счетчика
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value // Возвращаем текущее значение
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	// Запускаем 10 горутин для инкрементации счетчика
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait() // Ожидаем завершения всех горутин
	fmt.Println("Итоговое значение счетчика:", counter.Value())
}
