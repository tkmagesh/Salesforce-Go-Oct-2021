package main

import "fmt"

func main() {
	//Arrays
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	var nos [5]int = [...]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	//iteration using indexer
	fmt.Println("iteration using indexer")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(nos[idx])
	}

	//iteration using range
	fmt.Println("iteration using range")
	for _, value := range nos {
		fmt.Println(value)
	}

	//copy of an array
	var newNos [5]int
	newNos = nos
	fmt.Println(newNos)
	newNos[0] = 200
	fmt.Println(newNos, nos)
}
