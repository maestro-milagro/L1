package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными
//параметрами x,y и конструктором.

// Зададим структуру Point и зададим два поля типа float64 для x и y
type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Определим метод находящий расстояние между двумя точками как корень из суммы квадратов их сторон
func findDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))
}
func main() {
	x1 := 1.0
	x2 := 2.0
	y1 := 1.0
	y2 := 2.0
	p1 := NewPoint(x1, y1)
	p2 := NewPoint(x2, y2)
	fmt.Println(findDistance(*p1, *p2))
}
