package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1)
	b := big.NewInt(1)
	a.Lsh(a, 21) // a = 2^21
	b.Lsh(b, 22) // b = 2^22

	// Сложение
	sum := new(big.Int).Add(a, b)
	fmt.Println("Сумма:", sum)

	// Вычитание
	diff := new(big.Int).Sub(a, b)
	fmt.Println("Разность:", diff)

	// Умножение
	product := new(big.Int).Mul(a, b)
	fmt.Println("Произведение:", product)

	// Деление
	quotient := new(big.Int).Div(a, b)
	fmt.Println("Частное:", quotient)
}
