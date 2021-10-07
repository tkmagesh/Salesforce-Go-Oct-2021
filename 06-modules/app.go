package main

import (
	"fmt"
	calculator "modularity-demo/calc"
	"modularity-demo/calc/utils"

	"github.com/fatih/color"
)

func main() {
	color.Red(fmt.Sprintln(calculator.Add(1, 2)))
	color.Blue(fmt.Sprintln(calculator.Subtract(1, 2)))
	color.Green(fmt.Sprintln(calculator.Subtract(1, 2)))

	fmt.Println(utils.IsEven(100))
}
