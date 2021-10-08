//Share memory by communicating - ADVISABLE

//simulating a deadlock
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	ch := make(chan int)
	go add(100, 200, ch)
	wg.Wait()
	result := <-ch //reading the result from the channel
	fmt.Println(result)
	fmt.Println("Exiting main")
}

func add(x, y int, ch chan int) {
	time.Sleep(2 * time.Second)
	result := x + y
	ch <- result //writing the result to the channel
	wg.Done()
}

/*
ch <- value //writing data into the channel
value := <-ch //reading data from the channel
*/
