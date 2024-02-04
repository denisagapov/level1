package main

import (
	"fmt"
	"time"
)

// функция принимает продолжительность сна
func sleep(d time.Duration) {
	// time.After создает канал, который будет возвращает время, через
	// указанное количество времени, а оператор <- блокирует выполнение
	// до тех пора, пока не будет получено новое значение
	<-time.After(d)
}
func main() {
	fmt.Println("старт")
	// передаем в функцию sleep 5 секунд
	sleep(5 * time.Second)
	fmt.Println("стоп")
}
