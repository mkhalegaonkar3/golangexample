package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, "\t", p.age)
}
func main() {
	p1 := person{
		first: "Mayuresh",
		last:  "Khalegaonkar",
		age:   23,
	}
	fmt.Println("NAME\t\tAGE")
	p1.speak()
}
