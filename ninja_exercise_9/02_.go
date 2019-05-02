package main

import "fmt"

type person struct {
	fname string
	lname string
	age   int
}

func (p *person) speak() {
	fmt.Println("From speak method", p.lname)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		fname: "Mayuresh",
		lname: "Khalegaonkar",
		age:   23,
	}

	//saySomething(p1)
	/* .\02_.go:30:14: cannot use p1 (type person) as type human in argument to saySomething:
	person does not implement human (speak method has pointer receiver)*/
	saySomething(&p1)

}
