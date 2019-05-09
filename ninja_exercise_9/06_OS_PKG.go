package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println("Go OS =", runtime.GOOS)
	fmt.Println("Arch=", runtime.GOARCH)
	fmt.Println("Number of CPUs", runtime.NumCPU())
	fmt.Println("Number of go routines", runtime.NumGoroutine())
	fmt.Println("Number of go root", runtime.GOROOT)
}
