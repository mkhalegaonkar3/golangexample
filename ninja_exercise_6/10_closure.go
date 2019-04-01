package main

import "fmt"

func main() {
	f := display()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	fmt.Println("............Another scope")
	g := display()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

}
func display() func() int {
	c := 0
	return func() int {

		c++
		return c
	}

}
