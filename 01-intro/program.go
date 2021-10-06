package main

import "fmt"

//unused variables at package level are permitted (NOT AT FUNCTION LEVEL)
var no int

func main() {
	/*
		var str string
		str = "Hello, world!"
	*/

	/*
		var str string = "Hello, world!"
	*/

	/*
		//type inference
		var str = "Hello, world!"
	*/

	//the following syntax is applicable only in a function (not at package level)
	str := "Hello, world!"
	fmt.Println(str)

	//declaring multiple variables
	/*
		var x int
		var y int
		var msg string
		x = 100
		y = 200
		msg = "Result :"
	*/

	/*
		var x, y int
		var msg string
		x = 100
		y = 200
		msg = "Result :"
	*/

	/*
		var (
			x, y int
			msg  string
		)
		x = 100
		y = 200
		msg = "Result :"
	*/

	/*
		var (
			x   int    = 100
			y   int    = 200
			msg string = "Result :"
		)
	*/
	/*
		var (
			x, y int    = 100, 200
			msg  string = "Result :"
		)
	*/

	/*
		var (
			x, y = 100, 200
			msg  = "Result :"
		)
	*/

	/*
		var (
			x, y, msg = 100, 200, "Result :"
		)
	*/

	/*
		var x, y, msg = 100, 200, "Result :"
	*/

	x, y, msg := 100, 200, "Result :"
	fmt.Println(msg, x+y)

	//constants
	const pi = 3.14

	//iota
	/*
		const (
			red   = iota
			green = iota
			blue  = iota
		)
	*/

	/*
		const (
			red = iota
			green
			blue
		)
	*/

	/*
		const (
			red = iota + 2
			green
			blue
		)
	*/

	const (
		red = iota + 2
		green
		_
		blue
	)

	fmt.Println(red, green, blue)

	//type conversions (explicit)
	var a int = 100
	var b float32

	b = float32(a)
	fmt.Println(a, b)

	var myVar int
	myVar = 100
	fmt.Println("Value of myVar = ", myVar)

}
