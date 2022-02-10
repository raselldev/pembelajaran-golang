package main

import "fmt"

func fullName() (string, string) {
	return "Raka", "Rasell"
}

func main() {
	firstName, lastName := fullName()

	fmt.Println(firstName, lastName)
}
