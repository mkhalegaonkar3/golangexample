package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go fun1()
	go fun2()
	wg.Wait()
	fmt.Println("main exit")
}
func fun1() {

	fmt.Println("from function 1")
	wg.Done()
}
func fun2() {
	fmt.Println("from function 2")
	wg.Done()
}
