package main

import "fmt"

type customErr struct {
	err string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("here is the error %v", ce.err)
	//return fmt.Sprintf("here is the error %v",ce.(customErr).err)//assertion
}
func foo(e error) {
	fmt.Println("foo ran", e)
}
func main() {
	c1 := customErr{
		err: "Hi there this is error",
	}
	fmt.Println(c1.Error())

	foo(c1)
}
