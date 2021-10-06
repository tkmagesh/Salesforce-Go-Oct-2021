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

	//switch construct
	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 6 => "Not Bad"
		score 7 - 9 => "Good"
		score 10 => "Excellent"
		otherwise => "invalid score"
	*/

	score := 5
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6:
		fmt.Println("Not Bad")
	case 7, 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("invalid score")
	}

	newNo := 11
	switch {
	case newNo%2 == 0:
		fmt.Println(newNo, " is even")
	case newNo%2 == 1:
		fmt.Println(newNo, " is odd")
	}

	//no := 4
	var no int
	fmt.Print("Enter a number: ")
	fmt.Scanf("%d", &no)
	switch no {
	case 0:
		fmt.Println(" is zero")
	case 1:
		fmt.Println(" is <= 1")
		fallthrough
	case 2:
		fmt.Println(" is <= 2")
		fallthrough
	case 3:
		fmt.Println(" is <= 3")
		fallthrough
	case 4:
		fmt.Println(" is <= 4")
		fallthrough
	case 5:
		fmt.Println(" is <= 5")
		fallthrough
	case 6:
		fmt.Println(" is <= 6")
	case 7:
		fmt.Println(" is <= 7")
		fallthrough
	case 8:
		fmt.Println(" is <= 8")
	}
}
