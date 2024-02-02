package main

import "fmt"

// функция принимает переменную типа interface{} и возвращает тип данных
func getType(variable interface{}) {
	// switch сравнивает тип данных переменной в каждом case
	switch variable.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("channel int")
	default:
		fmt.Println("Другой тип переменной")
	}
}
func main() {
	var intNumber int = 2
	var boolType bool = true
	var stringType string = "строка"
	var channel chan int = make(chan int)

	getType(intNumber)
	getType(boolType)
	getType(stringType)
	getType(channel)
	getType(42.4)
}
