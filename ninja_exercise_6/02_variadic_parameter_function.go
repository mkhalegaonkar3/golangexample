package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5}
	n := sum(ii...) //unferling
	fmt.Println(n)

	ii2 := []int{1, 2, 3, 4, 5}
	n1 := bar(ii2) //passing slice
	fmt.Println(n1)
}

func sum(xi ...int) int {

	total := 0
	for _, v := range xi {

		total += v
	}
	return total
}

func bar(xi []int) int {

	total := 0
	for _, v := range xi {

		total += v
	}
	return total
}
