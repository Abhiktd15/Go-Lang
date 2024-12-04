package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the class of pointers!")

	// var ptr *int

	// fmt.Println("value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber
	//THis gives the address of the pointer
	fmt.Println("value of pointer is ", ptr)

	//THis give the refrence of the pointer(value )
	fmt.Println("value of pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("value of pointer after doubling is ", myNumber)
}
