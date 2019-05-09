package main

import (
	"fmt"
	"sync"
)

var counter int = 0
var wg sync.WaitGroup
var mux sync.Mutex

func main() {
	n := 100
	wg.Add(n)
	for i := 1; i <= n; i++ {

		go func() {
			mux.Lock()
			v := counter
			v = counter
			v++
			counter = v
			fmt.Println("value of incrementer", counter)
			mux.Unlock()
			wg.Done()
		}() //annonymous func

	}
	wg.Wait()
	fmt.Println("final counter", counter)
	fmt.Println("exited")
}
