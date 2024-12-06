package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the class of Structs in GoLang")

	//NO inheritance in golang; No Super or parent
	Abhishek := User{"Abhishek", "abhiktd15@gmail.com", true, 25}
	fmt.Println(Abhishek)
	fmt.Printf("Abhishek details are :%+v \n", Abhishek)
	fmt.Printf("Name is  :%v and email is %v \n", Abhishek.Name, Abhishek.Emails)
}

type User struct {
	Name   string
	Emails string
	Status bool
	Age    int
}
