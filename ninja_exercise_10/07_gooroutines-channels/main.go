package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {

		c := make(chan int)
		fmt.Println("go routine :", i)
		go func() {

			for i := 0; i < 10; i++ {
				c <- i

			}

			close(c)

		}()

		for v := range c {
			fmt.Println("\t\tchannel value : ", v)
		}
	}

	//fmt.Println(<-c)
	fmt.Println("About to exit")
}
