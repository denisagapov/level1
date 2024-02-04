package main

import (
	"fmt"
	"sync"
)

func squareNumbers(number int, wg *sync.WaitGroup) {
	// отложенный вызов метода Done для WaitGroup
	defer wg.Done()
	// вычисление квадрата для числа
	number *= number
	// вывод квадрата числа в stdout
	fmt.Println(number)
}

func main() {
	// создание объекта WaitGroup для синхронизации горутин
	var wg sync.WaitGroup
	// определение массива из 5 чисел
	numbers := [5]int{2, 4, 6, 8, 10}
	// цикл перебора всех элементов массива
	for _, number := range numbers {
		// увеличить счетчик WaitGroup на 1
		wg.Add(1)
		// запуск горутины
		go squareNumbers(number, &wg)
	}
	// ожидание завершения работы всех горутин
	wg.Wait()
}
