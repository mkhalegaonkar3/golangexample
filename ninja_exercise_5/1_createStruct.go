package main

import "fmt"

type person struct {
	fname   string
	lname   string
	hobbies []string
}

func main() {
	p1 := person{
		fname:   "mayuresh",
		lname:   "khalegaonkar",
		hobbies: []string{"cricket", "dancing"},
	}
	p2 := person{
		fname:   "swanand",
		lname:   "khalegaonkar",
		hobbies: []string{"painting", "surfing"},
	}
	sliceStruct := []person{p1, p2}
	//fmt.Printf("%T",sliceStruct)
	for _, v := range sliceStruct {

		fmt.Println("Record for ", v.fname)

		fmt.Println("\tFull Name :", v.fname, "", v.lname)
		for j, hobby := range v.hobbies {
			fmt.Println("\t\tHobby", j, hobby)
		}

	}

}
