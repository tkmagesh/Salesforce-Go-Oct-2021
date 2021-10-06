package main

import "fmt"

func main() {
	logAdd := logPerform(add)
	logSubtract := logPerform(subtract)

	logAdd(10, 20)
	logSubtract(100, 200)

}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func logPerform(op func(int, int) int) func(x, y int) {
	return func(x, y int) {
		fmt.Println("Before invocation:")
		result := op(x, y)
		fmt.Println(result)
		fmt.Println("After invocation:")
	}
}
