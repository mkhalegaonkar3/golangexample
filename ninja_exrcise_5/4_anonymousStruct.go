package main

import "fmt"

func main(){
	s1:=struct{
		fname string
		lname string
		marks []int

	}{
		fname:"mayuresh",
	    lname:"khalegaonkar",
	    marks: []int{121,123},
	}
	s2:=struct{
		fname string
		lname string
		marks []int
	}{
		fname:"swanand",
		lname:"khalegaonkar",
		marks: []int{123,1232},
	}

	m:=map [int][]string{
		1:{s1.fname,s1.lname},
		2:{s2.fname,s2.lname},
	}
	fmt.Println(m[1])
}