package main

import (
	"fmt"
	"time"
)

func main() {
	printString("Hello")
	printString("World")
	fmt.Println("Exiting main")
}

func printString(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
	}
}

/*
Desired output:
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/
