package main

import (
	"fmt"
)

func main(){

	m:=map[string][]string{

		"mayuresh":{"dancing","Art","tabla"},
		"swanand":{"painting","social work","friedship"},
		"vidya":{"singing","internet surfing","chating"},
	}
m["sanjay"]=[]string{"AA","SASD","ADA"}//adding key-value pair in MAP


	for k,v:=range m{

		fmt.Println("Hobbies for",k)
			for k,v2:=range v{
				fmt.Println("\t",k,v2)
			}

	}

	delete(m, "sanjay")//deleting key value pair from MAP
	fmt.Println(m)

}
