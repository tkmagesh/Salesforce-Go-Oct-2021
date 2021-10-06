package main

import "fmt"

func main() {
	//functions
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	var sum func(...int) int
	sum = func(nos ...int) int {
		var total int
		for _, v := range nos {
			total += v
		}
		return total
	}
	fmt.Println(sum(1, 2, 3, 4, 5))

	//anonymous functions
	func() {
		fmt.Println("anonymous function")
	}()

	func(x, y int) {
		fmt.Println(x, y)
	}(100, 200)

	quotient, remainder := func(x, y int) (int, int) {
		return x / y, x % y
	}(100, 7)
	fmt.Println(quotient, remainder)
}
