package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// В оригинальном коде подстрока ссылается на огромную строку, что расходует много памяти
	justString = string([]rune(v[:100])) // Создаем копию подстроки
}

func createHugeString(size int) string {
	hugeString := make([]rune, size)
	for i := 0; i < size; i++ {
		hugeString[i] = 'a' // Заполняем строку символами 'a'
	}
	return string(hugeString)
}

func main() {
	someFunc()
}
