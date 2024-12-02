package main

import "fmt"

const LogginToken string = "dlkfiehrjadadfa"

func main() {
	var username string = "abhishek"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var floatVal float32 = 255.324232342342343242
	fmt.Println(floatVal)
	fmt.Printf("Variable is of type: %T \n", floatVal)

	//default value and some aliases
	var anotherVariable int
	fmt.Println((anotherVariable))
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	//implicit type
	var website = "http://google.com"
	fmt.Println(website)

	//no var style
	no_of_user := 300000
	fmt.Println(no_of_user)

	fmt.Println(LogginToken)
	fmt.Printf("variable is of type: %T \n", LogginToken)

}
