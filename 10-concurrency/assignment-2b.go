package main

import "fmt"

/*
divide & conquer
	break the data in to two equal sets
	sum the data using 2 goroutintes
	get the sum results from the goroutines and print the final result
*/

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}
	dataSet1 := data[:len(data)/2]
	dataSet2 := data[len(data)/2:]
	ch := make(chan int)

	go func(ch chan int, nos []int) {
		ch <- sum(nos...)
	}(ch, dataSet1)

	go func(ch chan int, nos []int) {
		ch <- sum(nos...)
	}(ch, dataSet2)

	/*
		result1 := <-ch
		result2 := <-ch
		result := result1 + result2
	*/
	result := <-ch + <-ch
	fmt.Println(result)
}

func sum(nos ...int) int {
	var sum int
	for _, v := range nos {
		sum += v
	}
	return sum
}
