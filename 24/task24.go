package main

import (
	"fmt"
	"math"
)

// определяем структуру Point с полями x, y
type Point struct {
	x, y float64
}

// функция для определения конструктора для Point
func newPoint(x, y float64) *Point {
	return &Point{x, y}
}

// метод для получения значения x из структуры Point
func (p *Point) getX() float64 {
	return p.x
}

// метод для получения значения y из структуры Point
func (p *Point) getY() float64 {
	return p.y
}

// метод для получения расстояния между точками, с использованием пакета math
func (p *Point) getDistance(other *Point) float64 {
	dx := p.getX() - other.getX()   // разница по х
	dy := p.getY() - other.getY()   // разница по y
	return math.Sqrt(dx*dx + dy*dy) // расстояние между точками
}
func main() {
	p1 := newPoint(2.0, 3.0)
	p2 := newPoint(5.0, 7.0)
	fmt.Println(p1.getDistance(p2))
}
