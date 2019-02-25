package main

import (
	"fmt"
)

func main() {
	m1 := []string{"James", "Bond", "Shaken, not stirred"}
	m2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	multiSlice := [][]string{m1, m2}
	fmt.Println(multiSlice)

	for i, slice := range multiSlice {
		fmt.Println("Record", i)
		for j, v := range slice {
			fmt.Println("\t", j, v)
		}

	}
}
