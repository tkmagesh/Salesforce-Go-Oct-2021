package main

import "fmt"

type MyStr string

func (s MyStr) Length() int {
	return len(s)
}

func main() {
	//var s MyStr = "Hello, World!"
	s := MyStr("Hello, World!")
	fmt.Println(s.Length())
}
