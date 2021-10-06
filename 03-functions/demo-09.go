package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("[@main] deferred")
	}()
	f1()
	fmt.Println("main completed")
}

func f1() {
	defer func() {
		fmt.Println("[@f1] deferred - 1")
	}()
	defer func() {
		fmt.Println("[@f1] deferred - 2")
	}()
	defer func() {
		fmt.Println("[@f1] deferred - 3")
	}()
	f2()
	fmt.Println("f1 completed")
}

func f2() {
	defer func() {
		fmt.Println("[@f2] deferred - 1")
	}()
	defer func() {
		fmt.Println("[@f2] deferred - 2")
	}()
	defer func() {
		fmt.Println("[@f2] deferred - 3")
	}()
	fmt.Println("f2 completed")
}
