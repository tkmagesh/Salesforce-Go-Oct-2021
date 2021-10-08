package main

import (
	"fmt"
)

func main() {
	go f1()
	f2()
	//time.Sleep(time.Second * 2)
	/*
		var input string
		fmt.Scanln(&input)
	*/
	fmt.Println("Exiting main")
}

func f1() {

	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
