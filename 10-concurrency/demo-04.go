//Communicate by sharing memory

package main

import (
	"fmt"
	"sync"
)

var result int

var mutex sync.Mutex

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)
	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()
	fmt.Println(result)
}

func add(x, y int, wg *sync.WaitGroup) {
	mutex.Lock()
	{
		result = x + y
	}
	mutex.Unlock()
	wg.Done()
}
