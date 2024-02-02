package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// запускает воркера, который обрабатывает данные, получаемые из канала ch
// каждый воркер имеет свой уникальный id
func startWorkers(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // // отложенный вызов метода Done для WaitGroup
	for n := range ch {
		// выводит информацию о том, какой воркер работает с какими данными
		fmt.Printf("Воркер %d работает с данными №%d\n", id, n)
	}
}

func main() {
	var wg sync.WaitGroup // используется для ожидания завершения всех горутин
	var input string

	ch := make(chan int) // канал для передачи данных горутинам
	fmt.Println("Введите количество воркеров")
	fmt.Scanln(&input) // считывает количество воркеров как строку

	// преобразует строковое представление количества воркеров в число, для проверки, что было введено именно число
	numWorkers, err := strconv.Atoi(input)
	if err != nil {
		// выводит ошибку, если введенное значение не может быть преобразовано в число.
		fmt.Println("Ошибка: введите корректное целое число.")
		return

	}

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // добавляет счетчик горутин
		// запускает горутину воркера
		go startWorkers(i, ch, &wg)
	}

	signalCh := make(chan os.Signal, 1)     // канал для перехвата сигналов операционной системы
	signal.Notify(signalCh, syscall.SIGINT) // перехватывает сигнал (Ctrl+C)

	go func() {
		n := 0
		for true {
			select {
			// срабатывает, есть есть значение в канале signalCh
			case <-signalCh:
				// закрывает канал ch при получении сигнала прерывания, что приводит к завершению воркеров
				fmt.Println("Получен сигнал, завершаю работу")
				close(ch)
				return

			default:
				// отправляет данные в канал с задержкой, имитируя работу
				ch <- n
				n++
				time.Sleep(600 * time.Millisecond) // задержка между отправками данных
			}
		}
	}()
	wg.Wait() // ожидает завершения всех горутин
}
