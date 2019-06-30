package main

import (
	"fmt"

	"github.com/mkhalegaonkar3/golangexample/ninja_exercise_13/02_BET_on_given_code/quote"
	"github.com/mkhalegaonkar3/golangexample/ninja_exercise_13/02_BET_on_given_code/word"
)

func main() {

	fmt.Println("Word count is ", word.Count(quote.May))
	// for k, v := range word.UseCount(quote.May) {
	// 	fmt.Println(v, k)
	// }
}
