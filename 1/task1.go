package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) NameAge() {
	fmt.Printf("My name is %s, %d y.o.\n", h.Name, h.Age)
}

//		Определение структуры Action со встраиванием структуры Human
type Action struct {
	Human
}

func main() {
	action := Action{
		Human{Name: "Denis", Age: 21},
	}
	//		Использование методов из родительской структуры Human
	action.NameAge()       // Можно вызвать метод напрямую
	action.Human.NameAge() // Вызов метода на встроенном Human внутри action
}
