package main

import (
	"fmt"
)

// функция для переворачивания строки
func reverse(s string) string {
	// преобразуем строку в руны, т.к. они удобны для работы с unicode
	runes := []rune(s)
	// цикл берет начальное и последнее значение, меняет их местами
	// и сдвигается к середине, и с начала и с конца
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	// возвращаем преобразованную из рун строку
	return string(runes)
}
func main() {
	str := "главрыба"
	fmt.Println(reverse(str))
}
