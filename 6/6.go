package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Горутина, которая останавливается через канал
func workerWithChannel(stopChan <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopChan:
			fmt.Println("workerWithChannel: получен сигнал остановки")
			return
		default:
			fmt.Println("workerWithChannel: работа...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Горутина, которая останавливается через контекст
func workerWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("workerWithContext: получен сигнал остановки")
			return
		default:
			fmt.Println("workerWithContext: работа...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

// Горутина, которая останавливается через флаг (bool) с мьютексом
func workerWithFlag(stopFlag *bool, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		mutex.Lock()
		if *stopFlag {
			mutex.Unlock()
			fmt.Println("workerWithFlag: получен сигнал остановки")
			return
		}
		mutex.Unlock()
		fmt.Println("workerWithFlag: работа...")
		time.Sleep(500 * time.Millisecond)
	}
}

// Горутина, которая останавливается через таймер
func workerWithTimer(stopChan <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	timer := time.NewTimer(4 * time.Second)
	defer timer.Stop()
	for {
		select {
		case <-stopChan:
			fmt.Println("workerWithTimer: получен сигнал остановки")
			return
		case <-timer.C:
			fmt.Println("workerWithTimer: таймер истек")
			return
		default:
			fmt.Println("workerWithTimer: работа...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// Остановка горутины через канал
	stopChan := make(chan bool)
	wg.Add(1)
	go workerWithChannel(stopChan, &wg)
	time.Sleep(2 * time.Second)
	stopChan <- true

	// Остановка горутины через контекст
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go workerWithContext(ctx, &wg)
	time.Sleep(2 * time.Second)
	cancel()

	// Остановка горутины через флаг
	stopFlag := false
	var mutex sync.Mutex
	wg.Add(1)
	go workerWithFlag(&stopFlag, &mutex, &wg)
	time.Sleep(2 * time.Second)
	mutex.Lock()
	stopFlag = true
	mutex.Unlock()

	// Остановка горутины через таймер
	stopChanTimer := make(chan bool)
	wg.Add(1)
	go workerWithTimer(stopChanTimer, &wg)
	time.Sleep(3 * time.Second)
	stopChanTimer <- true

	wg.Wait()
	fmt.Println("Все воркеры остановлены.")
}
