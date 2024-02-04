// используем пакет math/big для работы с большими числами
// при передачи больших значений в функцию, используем указатели на числа
// чтобы избежать копирование больших чисел, что уменьшает производительность за счет потребления памяти
package main

import (
	"fmt"
	"math/big"
)

// принимает два указателя на big.Float и возвращает указатель на
// новый big.Float который равен сумме двух чисел
func add(a, b *big.Float) *big.Float {
	return new(big.Float).Add(a, b)
}

// принимает два указателя на big.Float и возвращает указатель на
// новый big.Float который равен разности двух чисел
func sub(a, b *big.Float) *big.Float {
	return new(big.Float).Sub(a, b)
}

// принимает два указателя на big.Float и возвращает указатель на
// новый big.Float который равен произведению двух чисел
func mul(a, b *big.Float) *big.Float {
	return new(big.Float).Mul(a, b)
}

// принимает два указателя на big.Float и возвращает указатель на
// новый big.Float который равен частному двух чисел
func quo(a, b *big.Float) *big.Float {
	return new(big.Float).Quo(a, b)
}

func main() {
	// создаем два новых числа big.Float, которые > 2^20 (>1048576)
	a := big.NewFloat(243943542)
	b := big.NewFloat(132943549)

	fmt.Println(add(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(mul(a, b))
	fmt.Println(quo(a, b))
}
