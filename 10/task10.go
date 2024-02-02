package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// Создаем карту для группировки температур
	groupTemperatures := make(map[int][]float64, len(temperatures))
	// Проходим по всем температурам
	for _, temp := range temperatures {
		// Определяем группу для текущей температуры
		// Округляем каждую температуру к ближайшей по модулю вниз и при этом сохраняем знак при вычислениях
		group := int(math.Copysign(math.Abs(temp)/10.0, temp)) * 10
		// Добавляем температуру в соответствующую группу
		groupTemperatures[group] = append(groupTemperatures[group], temp)
	}
	// Выводим результат
	fmt.Println(groupTemperatures)
}
