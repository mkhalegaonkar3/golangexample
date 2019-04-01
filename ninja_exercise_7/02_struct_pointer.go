package main

import "fmt"

type person struct {
	fname string
	lname string
}

func changeMe(p *person) {
	p.fname = "Swanand"

	//(*p).fname="nandu"
	//fmt.Printf("%T\n",p)

}
func main() {
	p := person{
		fname: "Mayuresh",
		lname: "Khalegaonkar",
	}
	fmt.Println(p) //before

	fmt.Printf("%T\n", p)
	changeMe(&p)

	fmt.Println(p.fname)
}
