package main

import (
	"fmt"

	dog "github.com/mkhalegaonkar3/golangexample/ninja_exercise_12/01_creating_dog_package"
)

func main() {

	humanYear := 5
	fmt.Println("Entered Human years = ", humanYear)
	fmt.Println("Dog years are =", dog.Years(humanYear))
}
