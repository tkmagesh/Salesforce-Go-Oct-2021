/*
Problem : Keep generating the fibonacci series till the user hits ENTER key
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan string)
	go fibonacci(ch, done)
	go func() {
		var input string
		fmt.Scanln(&input)
		done <- "done"
	}()
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Exiting from main")
}

func fibonacci(ch chan int, done chan string) {
	defer close(ch)
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
