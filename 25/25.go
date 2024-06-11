package main

import (
	"fmt"
	"time"
)

// Функция sleep приостанавливает выполнение программы на заданное количество секунд
func sleep(seconds int) {
	<-time.After(time.Duration(seconds) * time.Second)
}

func main() {
	fmt.Println("Начало")
	sleep(5) // Сон на 5 секунды
	fmt.Println("Конец")
}
