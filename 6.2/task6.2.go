package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// создаем контекст с функцией отмены
	ctx, cancel := context.WithCancel(context.Background())

	// запуск горутины с бесконечным циклом for
	go func() {
		for {
			select {
			case <-ctx.Done(): // получение сигнала о завершении контекста
				return // завершение горутины
			default:
				fmt.Println("Выполняется работа")

			}
		}
	}()

	time.Sleep(time.Second * 2) // имитация работы
	cancel()                    // отправка сигнала о завершении

}
