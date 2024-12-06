package main

import (
	"fmt"
)

func main() {
	fmt.Println("IF Else in GoLang")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10 login Count"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("9 is an even number")
	} else {
		fmt.Println("9 is an odd number")
	}

	//Special syntax
	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than or equal to 10")
	}
}
