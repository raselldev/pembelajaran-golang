package main

import "fmt"

func fullName() (firstName, lastName string) {
	firstName = "Raka"
	lastName = "Rasell"

	return
}

func main() {
	firstName, _ := fullName()

	fmt.Println(firstName)
}
