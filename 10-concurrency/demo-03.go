/* Introducing WaitGroup */

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go f1(wg)
	wg.Add(1)
	go f2(wg)
	wg.Wait() //=> blocked until the wg counter becomes zero
	fmt.Println("Exiting main!")
}

func f1(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("f1 invoked")
	wg.Done() // => decrement the wg counter
}

func f2(wg *sync.WaitGroup) {
	fmt.Println("f2 invoked")
	wg.Done() // => decrement the wg counter
}
