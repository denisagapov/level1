package main

import "fmt"

// Функция для пересечения двух множеств
func intersect(slice1, slice2 []int) []int {
	// Отображение для отслеживания всех элементов первого множества
	elements := make(map[int]bool)
	// Срез для хранения результата
	var result []int
	// Заполняем отображение элементами первого множества
	for _, elem := range slice1 {
		elements[elem] = true
	}
	// Поиск общих элементов
	for _, elem := range slice2 {
		if elements[elem] {
			result = append(result, elem)
		}
	}
	return result
}

func main() {
	// Два неупорядоченных множества в виде срезов
	set1 := []int{1, 2, 3, 4, 5, 6, 9, 10}
	set2 := []int{4, 5, 6, 7, 8}
	// Вывод результата
	fmt.Println(intersect(set1, set2))
}
