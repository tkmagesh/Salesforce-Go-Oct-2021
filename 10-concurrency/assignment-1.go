package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	ch1 := make(chan string, 1) //fix for the deadlock
	ch2 := make(chan string)
	go printString("Hello", ch1, ch2)
	go printString("World", ch2, ch1)
	ch1 <- "start"
	wg.Wait() // reports a deadlock. To be fixed!
	fmt.Println("Exiting main")
}

func printString(str string, in, out chan string) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(500 * time.Millisecond)
		fmt.Println(str)
		out <- "done"
	}
	wg.Done()
}

/*
Desired output:
Hello
World
Hello
World
Hello
World
Hello
World
Hello
World
*/
