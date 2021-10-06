package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int = &no
	fmt.Println(noPtr, no)

	//dereferencing (accessing the value using the pointer)
	var val int = *noPtr
	fmt.Println(val)

	// no == *(&no)
}
