package main

import (
	"fmt"
)

func main() {
	cs := make(chan<- int)
	//solution to this error is
	//cs := make(chan int) use bi-directional channel

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}
