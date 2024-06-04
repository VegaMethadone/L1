package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Pair struct {
	x int
	y int
}

func NewPair(x, y int) Pair {
	return Pair{x: x, y: y}
}

func calculateDistance(first, second Pair) float64 {
	xLine := max(first.x, second.x) - min(first.x, second.x)
	yLine := max(first.y, second.y) - min(first.y, second.y)
	distance := math.Sqrt((float64(xLine*xLine + yLine*yLine)))
	return distance
}

func main() {
	xTest := NewPair(6, 8)
	yTest := NewPair(3, 4)
	fmt.Println(calculateDistance(xTest, yTest))
}
