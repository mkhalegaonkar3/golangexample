package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func main() {
	n := 100
	wg.Add(n)
	for i := 0; i < n; i++ {

		go func() {
			//v := counter
			atomic.AddInt64(&counter, 1)
			fmt.Println("incrementor is ", atomic.LoadInt64(&counter))
			wg.Done()
		}() //annonymous func

	}
	wg.Wait()
	fmt.Println("final counter", counter)
	fmt.Println("exited")
}
