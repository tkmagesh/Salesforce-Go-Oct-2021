package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

type ShapeWithArea interface {
	Area() float64
}

type ShapeWithPerimeter interface {
	Perimeter() float64
}

/* type Shape interface {
	Area() float64
	Perimeter() float64
} */

//interface composition
type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func main() {
	c := Circle{5}
	//fmt.Println("Area :", c.Area())
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{5, 10}
	//fmt.Println("Area :", r.Area())
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)
}

//utility function
func PrintArea(shape ShapeWithArea) {
	//print the area of the given object
	fmt.Println("Area :", shape.Area())
}

func PrintPerimeter(shape ShapeWithPerimeter) {
	//print the perimeter of the given object
	fmt.Println("Perimeter :", shape.Perimeter())
}

func PrintShape(s Shape) {
	PrintArea(s)
	PrintPerimeter(s)
}
