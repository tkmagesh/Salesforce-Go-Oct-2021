package main

import "fmt"

type Entity struct {
}

func (e Entity) M1() {

}

type Contract interface {
	M1()
}

func main() {
	var x interface{}
	x = 100
	x = "Hello"
	x = true

	//x = 1000
	if no, success := x.(int); success {
		y := no + 100
		fmt.Println(x, y)
	} else {
		fmt.Println("Cannot convert to int")
	}

	//x = 1000
	//x = "Hello"
	//x = true
	//x = []int{}
	x = Entity{}
	switch val := x.(type) {
	case int:
		fmt.Println(val + 5000)
	case string:
		fmt.Println(val + " World")
	case bool:
		fmt.Println(!val)
	case []int:
		fmt.Println("list of integers")
	case Contract:
		fmt.Println("type implementing the Contract interface")
	case Entity:
		fmt.Println("Its an Entity")
	default:
		fmt.Println("Unknown type")
	}

}
