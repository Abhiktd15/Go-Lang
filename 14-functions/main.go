package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in go lang")
	// greeting()
	// result := add(3, 4)
	// fmt.Println("result is ", result)
	// proResult := proAdder(2, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)
	// fmt.Println("ProResult  of all numbers is ", proResult)
	proResult, proMEssage := proAdder(2, 2, 56, 7, 8)
	fmt.Println("ProResult of all numbers is ", proResult)
	fmt.Println("Promessage of all numbers is ", proMEssage)
}
func add(val1 int, val2 int) int {
	return val1 + val2
}

// func proAdder(values ...int) int {
// 	total := 0
// 	for _, v := range values {
// 		total += v
// 	}
// 	return total
// }
func proAdder(values ...int) (int, string) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, "Hello This is a proMessgae"
}
func greeting() {
	fmt.Printf("Namastey from golang \n")
}
