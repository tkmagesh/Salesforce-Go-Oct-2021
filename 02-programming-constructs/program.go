package main

import "fmt"

func main() {
	//if construct
	/*
		no := 21
		if no%2 == 0 {
			fmt.Println(no, " is even")
		} else {
			fmt.Println(no, " is odd")
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Println(no, " is even")
	} else {
		fmt.Println(no, " is odd")
	}
	//fmt.Println("Value of no is ", no) => 'no' is not accessible outside the if block

	//for construct
	sum := 1
	for idx := 0; idx < 10; idx++ {
		sum += sum
	}
	fmt.Println("Sum is ", sum)

	//for constuct as a while loop
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println("numSum is ", numSum)

	//indefinite loop
	x := 1
	for {
		if x > 100 {
			break
		}
		x += x
	}
	fmt.Println("x is ", x)

}
