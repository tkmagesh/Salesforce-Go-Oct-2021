/* Introducing WaitGroup */

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go f1()
	wg.Add(1)
	go f2()
	wg.Wait() //=> blocked until the wg counter becomes zero
	fmt.Println("Exiting main!")
}

func f1() {
	time.Sleep(3 * time.Second)
	fmt.Println("f1 invoked")
	wg.Done() // => decrement the wg counter
}

func f2() {
	fmt.Println("f2 invoked")
	wg.Done() // => decrement the wg counter
}
