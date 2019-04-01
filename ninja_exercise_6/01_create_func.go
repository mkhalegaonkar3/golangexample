package main

import "fmt"

func main() {

	func1 := foo()
	n, st := bar()
	fmt.Println(func1)
	fmt.Println(n, st)

}

func foo() int {

	return 43
}
func bar() (int, string) {

	return 32, "mayuresh"
}
