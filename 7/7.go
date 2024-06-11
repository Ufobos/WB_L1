package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	m := make(map[int]int)

	// Функция для записи данных
	writeToMap := func(key, value int) {
		defer wg.Done()
		mu.Lock() // Блокирование map для записи
		m[key] = value
		mu.Unlock() // Разблокирование map
	}

	// Запуск нескольких горутин для конкурентной записи
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writeToMap(i, i*i)
	}

	wg.Wait() // Ожидание завершения всех горутин

	for k, v := range m {
		fmt.Printf("key: %d, value: %d\n", k, v)
	}
}

// Также можно использовать sync.Map,
// которая обеспечивает потокобезопасный доступ к map
//
// func main() {
// 	var sm sync.Map
// 	var wg sync.WaitGroup
// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(n int) {
// 			defer wg.Done()
// 			for j := 0; j < 10; j++ {
// 				sm.Store(n*10+j, n*10+j)
// 			}
// 		}(i)
// 	}
