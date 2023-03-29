package main

import (
	"fmt"
	"math"
)

// Define interface

type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{5}
	rectangle := Rectangle{4, 5}

	fmt.Println("Area of circle:", math.Round(getArea(circle)*100)/100) // To 2 decimal points
	fmt.Println("Area of rectangle:", getArea(rectangle))
}
