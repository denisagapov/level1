package main

import (
	"fmt"
)

var justString string

// функция createHugeString создает строку размером size байт
func createHugeString(size int) string {
	return string(make([]byte, size)) // инициализирует срез байт заданного размера и преобразует его в строку
}

// функция someFunc создает большую строку и сохраняет подстроку в переменную justString
func someFunc() {
	v := createHugeString(1 << 10)       // создает строку размером 1024 байта
	justString = string([]byte(v[:100])) // преобразует первые 100 байтов строки v в новую строку
	fmt.Println(len(justString))         // выводит длину новой строки
}
func main() {
	someFunc()
}
