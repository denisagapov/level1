package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()         // блокируем мьютекс, чтобы доступ к счетчику одновременно был только у 1 горутины
	defer c.mu.Unlock() // разблокировка мьютекса после выхода из функции
	c.value++
}

func main() {
	counter := Counter{value: 0} // создаем базовое значение счетчика
	var wg sync.WaitGroup
	numGo := 10
	wg.Add(numGo) // добавляем счетчик горутин, равным 10
	// запускаем горутины, выполняющие метод Increment
	for i := 0; i < numGo; i++ {
		go func() {
			defer wg.Done() // после выхода из функции уменьшаем счетчик горутин на 1
			counter.Increment()
		}()
	}
	wg.Wait() // ожидание завершения работы все горутин
	fmt.Println("Значение счетчика:", counter.value)
}
