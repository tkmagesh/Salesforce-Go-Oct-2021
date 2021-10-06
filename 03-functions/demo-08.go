package main

import "fmt"

func main() {
	increment := getCounter()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
}

func getCounter() func() int {
	var counter int = 0 // => closure
	return func() int {
		counter++
		return counter
	}
}
