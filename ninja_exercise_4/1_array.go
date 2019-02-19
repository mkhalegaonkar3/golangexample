package main

import (
	"fmt"
)

func main() {
	intArray := []int{12, 23, 43, 21, 332}
	stringArray:=[]string{"Mayuresh","swanand","Mayuresh","swanand","Mayuresh"}
	fmt.Println("Array representation",intArray )

	fmt.Println("by using for..range")
	fmt.Println("INDEX\tVALUE")
	for i,v :=range intArray{

		fmt.Println(i,"\t",v)
	}
	for i,v :=range stringArray{

		fmt.Println(i,"\t",v)
	}

	fmt.Printf("%T%T",intArray,stringArray)
	//fmt.Println(cap(intArray))
	//fmt.Println(len(stringArray))
}
