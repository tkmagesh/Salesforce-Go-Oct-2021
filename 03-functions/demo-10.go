package main

import (
	"fmt"
)

func main() {

	if result, err := divide(100, 0); err != nil {
		fmt.Println(err)
		//debug.PrintStack()
		return
	} else {
		fmt.Println(result)
	}
}

func divide(x, y int) (result int, err error) {
	defer func() {
		var er = recover()
		fmt.Println("er : ", er)
		var er2 = recover()
		fmt.Println("er2 : ", er2)
		if er != nil {
			//fmt.Println("Something went wrong : ", er)
			err = fmt.Errorf("%s", er)
		}

	}()
	if y == 0 {
		panic("Cannot divide by 0")
	}
	return x / y, nil
}
