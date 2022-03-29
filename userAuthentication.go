package main

import "fmt"

func main() {

	var userInputName string

	fmt.Println("Please key in your username:")
	_, _ = fmt.Scanln(&userInputName)

	if userInputName == "Admin" {
		fmt.Println("Welcome, Admin!")
	} else if userInputName == "Robin" {
		fmt.Println("Welcome!", userInputName)
	} else if userInputName == "John" {
		fmt.Println("Welcome!", userInputName)
	} else {
		fmt.Println("Merry Men")
	}

}
