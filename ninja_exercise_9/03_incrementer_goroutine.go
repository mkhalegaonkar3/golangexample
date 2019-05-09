package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0
var wg sync.WaitGroup

func main() {
	n := 100
	wg.Add(n)
	for i := 1; i <= n; i++ {

		go func() {
			v := counter
			runtime.Gosched()
			v = counter
			v++
			counter = v
			wg.Done()
		}() //annonymous func
		fmt.Println("value of incrementer", counter)

	}
	wg.Wait()
	fmt.Println("final counter", counter)
	fmt.Println("exited")
}
