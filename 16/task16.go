package main

import "fmt"

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

// функция быстрой сортировки, выполняется рекурсивно до тех пора, пока подмассивы не уменьшатся до 1 элемента
func quickSort(arr []int, low, high int) {
	if low < high {
		// находим индекс опорного элемента
		pivotIndex := partition(arr, low, high)
		// осуществляем быструю сортировку левой части массива
		quickSort(arr, low, pivotIndex-1)
		// осуществляем быструю сортировку правой части массива
		quickSort(arr, pivotIndex+1, high)
	}
}

func partition(arr []int, low, high int) int {
	// выбираем последний элемент как опорный
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		// если элемент меньше или равен опорному
		if arr[j] <= pivot {
			i++
			// происходит обмен местами элементов
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// опорный элемент помещается после всех чисел, которые меньше его
	arr[i+1], arr[high] = arr[high], arr[i+1]
	// возвращаем новую позицию опорного элемента
	return i + 1
}
