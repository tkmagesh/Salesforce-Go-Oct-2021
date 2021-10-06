package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func main() {
	var no int = 10
	var noPtr *int = &no
	fmt.Println(noPtr, no)

	//dereferencing (accessing the value using the pointer)
	var val int = *noPtr
	fmt.Println(val)

	// no == *(&no)
	fmt.Println("Before incrementing, no : ", no) //=> 10
	increment(&no)
	fmt.Println("After incrementing, no : ", no) //=> 11

	x, y := 10, 20
	fmt.Printf("Before swapping x = %d and y = %d\n", x, y) //=> 10, 20
	swap(&x, &y)
	fmt.Printf("After swapping x = %d and y = %d\n", x, y) //=> 20 ,10
}

func increment(noPtr *int) {
	*noPtr++
}

func swap(x, y *int) {
	/*
		var temp int
		temp = *x
		*x = *y
		*y = temp
	*/
	*x, *y = *y, *x
}
