package main

import "fmt"

// определяем интерфейс Animal с методом MakeSound
type Animal interface {
	MakeSound() string
}

// структура для собаки
type Dog struct{}

// собака реализует метод MakeSound (подходит интерфейсу Animal)
func (d *Dog) MakeSound() string {
	return "Woof"
}

// структура для кошки
type Cat struct{}

// кошка реализует метод Meow (не подходит интерфейсу Animal)
func (c *Cat) Meow() string {
	return "Meow"
}

// структура адаптера, которая принимает указатель на структуру кошки
type Adapter struct {
	cat *Cat
}

// адаптер реализует метод MakeSound (подходит интерфейсу Animal)
// он возвращает метод Meow у кошки -> теперь кошка удовлетворяет интерфейсу
func (a *Adapter) MakeSound() string {
	return a.cat.Meow()
}

func main() {
	dog := &Dog{}                    // определяем собаку
	cat := &Cat{}                    // определяем кошку
	adapter := &Adapter{cat: cat}    // создаем адаптер и передаем в него кошку
	fmt.Println(dog.MakeSound())     // выводит звук собаки
	fmt.Println(adapter.MakeSound()) // выводит звук кошки через адаптер
}
