package main

import (
	"fmt"
	"time"
)

func main() {
	durable := 5 // время работы программы
	ch := make(chan int)

	// горутина, которая запуска бесконечный цикл для отправки данных в канал
	go func() {
		for i := 0; ; i++ {
			// отправка значения i в канал ch
			ch <- i
		}
	}()

	// горутина, которая считывает все данные из канала
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	// остановка программы по истечению времени
	time.Sleep(time.Duration(durable) * time.Second)
	fmt.Println("Время вышло")
}
