package main

import (
	"fmt"
	calculator "modularity-demo/calc"
	"modularity-demo/calc/utils"
)

func main() {
	fmt.Println(calculator.Add(1, 2))
	fmt.Println(calculator.Subtract(1, 2))
	fmt.Println(calculator.Subtract(1, 2))
	fmt.Println(calculator.Subtract(1, 2))

	fmt.Println(calculator.GetOpCount())

	fmt.Println(utils.IsEven(100))
}
