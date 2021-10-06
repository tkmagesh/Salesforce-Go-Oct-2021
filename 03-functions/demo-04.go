package main

import "fmt"

func main() {
	//exec(fn)
	/* exec(func() {
		fmt.Println("Hello")
	}) */

	/*
		fmt.Println("Before invocation:")
		fmt.Println(add(100, 200))
		fmt.Println("After invocation:")

		fmt.Println("Before invocation:")
		fmt.Println(subtract(100, 200))
		fmt.Println("After invocation:")
	*/

	logPerform(add, 100, 200)
	logPerform(subtract, 100, 200)
}

/* func fn() {
	fmt.Println("fn is invoked")
} */

func exec(f func()) {
	fmt.Println("Executing")
	f()
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func logPerform(op func(int, int) int, x, y int) {
	fmt.Println("Before invocation:")
	result := op(x, y)
	fmt.Println(result)
	fmt.Println("After invocation:")
}
