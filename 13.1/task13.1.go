package main

import "fmt"

var (
	number1 = 20
	number2 = 10
)

func main() {
	fmt.Println("Числа: ", number1, number2)
	number1, number2 = number2, number1
	fmt.Println("После перестановки: ", number1, number2)

}
