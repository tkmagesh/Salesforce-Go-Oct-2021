package main

import "fmt"

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
}
