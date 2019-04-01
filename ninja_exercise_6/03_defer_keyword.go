package main

import "fmt"

func main() {

	defer foo() /*we called this function first but deu to defer it will execute after bar()function excutes*/
	bar()
}

func foo() {
	fmt.Println("This is from foo :1")
}
func bar() {
	fmt.Println("This is from bar :2")
}
