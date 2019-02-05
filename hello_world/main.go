package main

import "fmt"

func main() {
	fmt.Println("Hi this is my program in golang.")

	foo1()

	for i := 0; i < 100; i++ {
		if i%2 != 0 {
			fmt.Println("loading...")
			fmt.Println(i)
		}
	}
	bar1()
}

func foo1() {
	fmt.Println("This is my first function in golang , welcome to foo world ")
}
func bar1() {
	fmt.Println("thanks for bearing us through out the program....")
}
