package main

import (
	"fmt"
)

// setBit устанавливает i-й бит в value на 1 или 0 в зависимости от параметра bit
func setBit(value int64, i int, bit bool) int64 {
	if bit {
		// Установка i-го бита в 1
		return value | (1 << i)
	} else {
		// Установка i-го бита в 0
		return value &^ (1 << i)
	}
}

// 1010

func main() {
	var value int64 = 10 // 1010 в двоичном виде
	fmt.Printf("Начальное значение: %4b\n", value)

	// Устанавливаем 2-й бит в 1
	value = setBit(value, 2, true)
	fmt.Printf("После установки 3-го бита в 1: %4b\n", value)

	// Устанавливаем 1-й бит в 0
	value = setBit(value, 1, false)
	fmt.Printf("После установки 1-го бита в 0: %4b\n", value)
}
