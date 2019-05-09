package main

import "fmt"

func main() {

	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	//OR
	// ch := make(chan int, 1)
	// ch <- 42

	fmt.Println(<-ch)
}
