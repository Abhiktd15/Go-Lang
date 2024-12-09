package main

import "fmt"

func main() {
	fmt.Println("Welcome to class of methods in GO lang")
	abhishek := User{"abhishek", "abhiktd15@gmail.com", true, 20}
	fmt.Printf("Abhishek details are : %v \n", abhishek)
	abhishek.getStatus()
	abhishek.newEmail()
	fmt.Printf("Email of the user is : %v \n", abhishek.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("Is user Active:", u.Status)

}
func (u User) newEmail() {
	u.Email = "test@example.com"
	fmt.Println("Email of the user is :", u.Email)
}
