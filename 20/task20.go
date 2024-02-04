package main

import (
	"fmt"
	"strings"
)

// функция для переворачивания строки
func reverseWords(s string) string {
	// разделяем строку на слова
	words := strings.Fields(s)
	// цикл проходит по всем словам в срезе и меняет их местами
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	// объединяем слова в строку
	return strings.Join(words, " ")
}
func main() {
	str := "snow dog sun"
	fmt.Println(reverseWords(str))
}
