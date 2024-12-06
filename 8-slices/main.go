package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the classses of Slices!")

	// //syntax to initialize slices
	// var fruitlist = []string{"Apple", "Banana", "Orange"}
	// fmt.Printf("Types of fruit list is %T \n", fruitlist)

	// fruitlist = append(fruitlist, "Mango", "Peach")

	// fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[1:3])

	// fmt.Println(fruitlist)

	// //ANother syntac walrus

	// //This make operator allocates and initializes the value
	// highScore := make([]int, 4)
	// highScore[0] = 234
	// highScore[1] = 345
	// highScore[2] = 456
	// highScore[3] = 567

	// //using append reallocates the memory in highscore
	// highScore = append(highScore, 555, 666, 777)

	// fmt.Println(highScore)
	// fmt.Println(len(highScore))

	// //This package sorts the high scores

	// sort.Ints(highScore)
	// fmt.Println(highScore)

	//How to remove the value from slice based on index
	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
