package main

import "fmt"

// setBitOne устанавливает i-й бит в 1
func setBitOne(number int64, i uint) int64 {
	number = number | (1 << i)
	return number
}

// setBitZero сбрасывает i-й бит в 0
func setBitZero(number int64, i uint) int64 {
	number = number & ^(1 << i)
	return number
}

func main() {
	var number int64
	var i uint

	// Ввод числа в десятичной системе счисления
	fmt.Println("Введите число в десятичной системе счисления:")
	fmt.Scan(&number)

	// Ввод номера бита, который вы хотите изменить
	fmt.Println("Введите номер бита, который вы хотите изменить:")
	fmt.Scan(&i)

	// Установка i-го бита в 1 и вывод результата
	fmt.Printf("%d бит изменился на 1 - %d\n", i, setBitOne(number, i-1))

	// Установка i-го бита в 0 и вывод результата
	setBitZero(number, i-1)
	fmt.Printf("%d бит изменился на 0 - %d \n", i, setBitZero(number, i-1))
}
