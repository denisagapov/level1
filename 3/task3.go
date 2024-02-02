package main

import (
	"fmt"
	"sync"
)

// вычисление квадрата числа и отправка результата в канал
func sumSquares(number int, ch chan int, wg *sync.WaitGroup) {
	// отложенный вызов метода Done для WaitGroup
	defer wg.Done()
	// отправка в канал квадрата числа
	ch <- number * number
}

func main() {
	var wg sync.WaitGroup
	numbers := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int, len(numbers)) // канал с буфером, который равен длине масива
	sum := 0

	// цикл для запуска горутины для каждого числа в массиве
	for _, num := range numbers {
		// увеличить счетчик WaitGroup на 1
		wg.Add(1)
		// запуск горутины
		go sumSquares(num, ch, &wg)
	}
	// ожидание завершения всех горутин
	wg.Wait()
	// закрытие канала после завершения всех горутин
	close(ch)

	// цикл для суммирование всех значений из канала
	for square := range ch {
		sum += square
	}

	fmt.Printf("Сумма квадратов: %d", sum)
}
