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
	var choice int
	var num1, num2 int
	var result int

	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		if choice == 5 {
			break
		}

		if choice < 1 || choice > 5 {
			fmt.Println("Invalid choice")
			continue
		}

		fmt.Print("Enter the first number: ")
		fmt.Scan(&num1)

		fmt.Print("Enter the second number: ")
		fmt.Scan(&num2)

		switch choice {
		case 1:
			result = num1 + num2
			fmt.Println("Result: ", result)
		case 2:
			result = num1 - num2
			fmt.Println("Result: ", result)
		case 3:
			result = num1 * num2
			fmt.Println("Result: ", result)
		case 4:
			result = num1 / num2
			fmt.Println("Result: ", result)
		}
	}
}
