//Communicate by sharing memory - NOT ADVISABLE

package main

import (
	"fmt"
	"sync"
)

type Result struct {
	value int
	sync.Mutex
}

func (r Result) SetValue(val int) {
	r.Lock()
	{
		r.value = val
	}
	r.Unlock()
}

var result Result

func main() {

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, wg)
	wg.Add(1)
	go add(100, 200, wg)
	wg.Wait()
	fmt.Println(result.value)
}

func add(x, y int, wg *sync.WaitGroup) {
	result.SetValue(x + y)
	wg.Done()
}
