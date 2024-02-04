package main

import "fmt"

func main() {
	numbers := [5]int{2, 4, 6, 8, 10}
	// создаем буферизированный канал с вместимостью=длине массива
	ch := make(chan int, len(numbers))
	// цикл перебора всех элементов массива
	for _, number := range numbers {
		// создаем горутины для каждого числа массива
		go func(num int) {
			// отправляе в канал квадрат числа
			ch <- (num * num)
		}(number)
	}
	// читаем из канала все значения
	for i := 0; i < len(numbers); i++ {
		fmt.Println(<-ch)
	}
}
