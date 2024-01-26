package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}

func FindDistance(point1, point2 Point) float64 {
	return math.Sqrt(math.Pow(float64(point1.X)-float64(point2.X), 2) + math.Pow(float64(point1.Y)-float64(point2.Y), 2))
}

func main() {
	point1 := NewPoint(1, 1)
	point2 := NewPoint(4, 7)
	fmt.Println(FindDistance(*point1, *point2))
}
