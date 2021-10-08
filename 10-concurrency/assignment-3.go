/* Using the "select" construct */
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(ch)

	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Exiting from main")
}

func fibonacci(ch chan int) {
	defer close(ch)
	done := time.After(20 * time.Second)
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			x, y = y, x+y
			time.Sleep(time.Millisecond * 500)
		case <-done:
			fmt.Println("Exiting from fibonacci")
			return
		}
	}

}
