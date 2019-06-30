package main

import (
	"fmt"

	dog1 "github.com/mkhalegaonkar3/golangexample/ninja_exercise_13/01_dog_pkg_BET"
)

func main() {

	humanYear := 5
	fmt.Println("Entered Human years = ", humanYear)
	fmt.Println("Dog years are =", dog1.Years(humanYear))
}
