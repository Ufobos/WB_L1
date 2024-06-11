package main

import (
	"fmt"
	"math"
)

// Функция для объединения температур в группы
func groupTemperatures(temperatures []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		key := int(math.Floor(temp/10) * 10)    // Создание ключа (с округлением)
		groups[key] = append(groups[key], temp) // Добавление значений
	}

	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := groupTemperatures(temperatures)

	for key, temps := range groups {
		fmt.Printf("%d: %v\n", key, temps)
	}
}
