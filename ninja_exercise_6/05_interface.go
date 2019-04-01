package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}
type circle struct {
	radious float64
}

func (c circle) area() float64 {
	return math.Pi * c.radious * c.radious
}
func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	circ := circle{
		radious: 12.343,
	}
	squa := square{
		length: 13,
	}
	info(circ)
	info(squa)

}
