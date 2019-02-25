package main

import "fmt"

func main() {

	intArray := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range intArray {

		fmt.Println(i, v)
	}
	intSlice := make([]int, 4)

	fmt.Println(intSlice)
	intSlice = append(intSlice, intArray...)
	fmt.Println(cap(intSlice))
	fmt.Println(cap(intArray))

	fmt.Println(intArray[:5])
	fmt.Println(intArray[5:])
	fmt.Println(intArray[2:7])
	fmt.Println(intArray[1:6])
	intSlice = append(intArray[:2], intArray[4:]...) //deleting slice
	fmt.Println(intSlice)

}
