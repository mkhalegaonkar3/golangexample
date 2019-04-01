package main

import "fmt"

var g func()

func main() {

	x := func() {
		fmt.Println("Assigned a function to a variable")
	}
	x() //function calling
	//fmt.Printf("%T\n",x)
	g = x //assigning a function variable
	g()
	fmt.Println("..............")
	fmt.Println("This is a anonymous function...")

	//anonymous function
	func() {

		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

	}() //anonymous function calling
}
