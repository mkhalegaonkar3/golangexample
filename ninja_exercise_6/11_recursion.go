package main

import "fmt"

func main() {
	n := 2
	result := fact(n)
	fmt.Println("Factorial of", n, "is", result)

}

func fact(n int) int {
	result := 1
	if n >= 1 {
		result = n * fact(n-1)
	}
	return result
}
