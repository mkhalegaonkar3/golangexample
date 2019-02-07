package main

import "fmt"

var y ="Mayuresh"
var z string =`Hi,"Good morning"`
func main(){
	fmt.Println(y)//var means global scope
	y:=34 //local scope inside function only
	fmt.Println(y) //it will print local value
	fmt.Printf("local variable is a type of %T\n",y)//it will print Type of the variable

	fmt.Println(z)
	fmt.Printf("%T\n",z)
	bar()
}

func bar(){

	fmt.Println(y)//it will print global value of y
	fmt.Printf(`"From Bar function "global varibale is a type of %T\n`,y)


}