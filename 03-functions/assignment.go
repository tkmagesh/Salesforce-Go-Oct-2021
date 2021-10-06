/*
	Create a console calculator which will display the following menu until the user choose to exit

	1. Add
	2. Subtract
	3. Multiply
	4. Divide
	5. Exit

	If the user chooses 1 - 4, accept two numbers from the user
	and perform the operation on them and print the result. Display the menu again

	if the user chooses any value but 1 - 4, then display the message "Invalid choice" and display the menu again
*/

package main

import "fmt"

func main() {
	for {
		choice := getMenuChoice()
		if choice == 5 {
			break
		}

		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		process(choice)
	}
}

func process(choice int) {
	num1, num2 := getOperands()
	oper := getOperation(choice)
	result := oper(num1, num2)
	fmt.Println("Result: ", result)
}

func getMenuChoice() int {
	var choice int
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")

	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)
	return choice
}

func getOperands() (int, int) {
	var num1, num2 int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)
	return num1, num2
}

func getOperation(choice int) (op func(int, int) int) {
	switch choice {
	case 1:
		op = add
	case 2:
		op = subtract
	case 3:
		op = multiply
	case 4:
		op = divide
	}
	return
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
