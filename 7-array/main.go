package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[3] = "Grapes"
	fmt.Println("Fruit list is :", len(fruitlist))

	var vegList = [3]string{"potato", "tomato", "mushrooms"}
	fmt.Println("Veg List is :", len(vegList))
}
