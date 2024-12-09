package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of loops in GOlang")

	// Array Loop
	// days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// fmt.Println(days)

	// For Loop
	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	//ANother type of loop
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is : %v and day is : %v \n", index, day)
	// }

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 5 {
			goto lco
		}
		fmt.Println("Rouge Value : ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at Abhishek.com")
}
