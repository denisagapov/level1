package main

import "fmt"

// функция удаляет элемент с помощью объединения всех элементов до
// указанного индекса и после
func delFormSlice(slice []int, i int) []int {
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}

// функция удаляет элемент с помощью перемещения последнего элемента слайса
// на место элемента с указанным индексом и удаляет последний элемент
func delFormSliceWithLastElement(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	slice := []int{2, 3, 5, 8, 9}
	index := 3
	fmt.Println(delFormSlice(slice, index))
	// fmt.Println(delFormSliceWithLastElement(slice, index))
}
