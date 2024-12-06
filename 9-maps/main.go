package main

import (
	"fmt"
)

func main() {
	fmt.Println("MAPS in GOlang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["TS"] = "TypeScript"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS Stands for : ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	//LOOPS in GOLANG
	//Loops in MAPS

	for key, value := range languages {
		fmt.Printf("For key: %v, value: %v\n", key, value)
	}

}
