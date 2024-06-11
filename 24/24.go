package main

import (
	"fmt"
	"math"
)

// Структура Point с параметрами x и y
type Point struct {
	x, y float64
}

// Конструктор для создания новой точки
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для вычисления расстояния между точками (Евклидово расстояние)
func (p *Point) DistanceTo(other *Point) float64 {
	return math.Sqrt(math.Pow(other.x-p.x, 2) + math.Pow(other.y-p.y, 2))
}

func main() {
	p1 := NewPoint(1, 2)
	p2 := NewPoint(4, 6)

	fmt.Printf("Расстояние между точками: %.2f\n", p1.DistanceTo(p2))
}
