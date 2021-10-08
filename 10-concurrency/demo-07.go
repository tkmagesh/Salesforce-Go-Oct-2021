//Share memory by communicating - ADVISABLE

package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int)
	go add(100, 200, ch)
	time.Sleep(5 * time.Second)
	result := <-ch //reading the result from the channel
	fmt.Println(result)
	fmt.Println("Exiting main")
}

func add(x, y int, ch chan int) {
	result := x + y
	fmt.Println("attempting to write the result to the channel")
	ch <- result //writing the result to the channel
	fmt.Println("result written to the channel")
}

/*
ch <- value //writing data into the channel
value := <-ch //reading data from the channel
*/
