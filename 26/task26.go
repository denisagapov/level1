package main

import (
	"fmt"
	"strings"
)

// функция принимает строку и возвращает булевое значение
func testUniq(str string) bool {
	// делаем строку в нижнем регистре
	str = strings.ToLower(str)
	// создаем мапу с ключом типа rune и значением типа bool
	m := make(map[rune]bool)
	// перебираем все символы строки
	for _, s := range str {
		// проверка, если символ уже есть в мапе (т.е. m[s]==true)
		if m[s] {
			return false // возвращаем false, если символ встретился
		}
		m[s] = true // добавляем символ в мапу с значением true
	}
	return true // возвращаем если повторений нет
}
func main() {
	str1 := "abcd"
	str2 := "abCdEaFeA"
	fmt.Println(str1, testUniq(str1)) // выведет true (нет повторов)
	fmt.Println(str2, testUniq(str2)) // выведет false (есть повторы)
}
