package main

import (
	"fmt"
)

func main() {
	n := 5
	a := 1

	for i := 0; i < n; i++ {
		for k := 5; k > i; k-- {
			fmt.Print(" ")

		}
		for j := 0; j <= i; j++ {
			fmt.Print(" ", a)
			a++
		}

		fmt.Println()
	}

}
