package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// создаем контекст с функцией отмены по таймеру
	ctx, _ := context.WithTimeout(context.Background(), time.Second*2)

	// запуск горутины с бесконечным циклом for
	go func() {
		for {
			select {
			case <-ctx.Done(): // получение сигнала о завершении контекста
				fmt.Println("Горутина завершилась")
				return // завершение горутины
			default:
				fmt.Println("Выполняется работа")

			}
		}
	}()
	time.Sleep(time.Second * 4) // установка таймера, чтобы основная горутина не завершалась до выполнения второстепенной
}
