package main

import "fmt"

func foo() func() {
	return func() {
		fmt.Println("function from foo")
	}

}
func bar(f func()) int {
	return 23
}
func main() {

	b := bar(foo()) //callback function
	foo()()
	fmt.Println("function bar", b)
}
