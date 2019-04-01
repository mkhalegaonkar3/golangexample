package main

import "fmt"

func add(a int, b int) func() int {
	return func() int {
		return a * b
	}

}
func main() {
	x := add(12, 3)
	fmt.Println(x())
}
