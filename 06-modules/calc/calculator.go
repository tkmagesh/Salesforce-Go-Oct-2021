package calc

import "fmt"

var opCount int

func init() {
	fmt.Println("calc.init()")
}

func GetOpCount() int {
	return opCount
}
