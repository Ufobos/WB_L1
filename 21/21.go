package main

import "fmt"

// Интерфейс, ожидаемый клиентом
type Target interface {
	Request() string
}

// Существующий класс с несовместимым интерфейсом
type Adaptee struct{}

func (a *Adaptee) SpecificRequest() string {
	return "Specific request"
}

// Адаптер, который позволяет Adaptee работать с интерфейсом Target
type Adapter struct {
	Adaptee *Adaptee
}

func (a *Adapter) Request() string {
	return a.Adaptee.SpecificRequest()
}

// Клиентская функция, использующая интерфейс Target
func clientCode(t Target) {
	fmt.Println(t.Request())
}

func main() {
	adaptee := &Adaptee{}
	adapter := &Adapter{Adaptee: adaptee}
	clientCode(adapter) // Вывод: Specific request
}
