package main

import "fmt"

func main(){
	fmt.Println("Println is a function /identifier of fmt package...it returns int n and error err ")
	n,err :=fmt.Println("Mayuresh")
	fmt.Println("you have typed",n,"characters with null character at the end of the every string")
	fmt.Println(err)
}
