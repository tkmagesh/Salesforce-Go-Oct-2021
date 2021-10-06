package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("Cannot divide by zero")

func main() {
	result, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("y cannot be zero")
		return
	}
	if err != nil {
		fmt.Println("something went wrong!", err)
		return
	}
	fmt.Println(result)
}

/*
func divide(x, y int) (int, error) {
	if y == 0 {
		return 0, DivideByZeroError
	}
	return x / y, nil
}
*/

func divide(x, y int) (result int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	result = x / y
	return
}
