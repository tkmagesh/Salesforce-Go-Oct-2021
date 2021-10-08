package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(ch)
	/* for i := 0; i < 15; i++ {
		fmt.Println(<-ch)
	} */
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Exiting from main")
}

func fibonacci(ch chan int) {
	defer close(ch)
	x, y := 0, 1
	for i := 0; i < 20; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- x
		x, y = y, x+y
	}

}
