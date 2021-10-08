//Share memory by communicating - ADVISABLE

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
	result := <-ch
	fmt.Println(result)
	wg.Wait()
	fmt.Println("Exiting main")
}

func add(x, y int, ch chan int) {
	time.Sleep(2 * time.Second)
	result := x + y
	ch <- result
	wg.Done()
}

/*
ch <- value //writing data into the channel
value := <-ch //reading data from the channel
*/
