package main

import "fmt"

func main() {
	name := "raka"

	switch name {
	case "Raka":
		fmt.Println("Hello Raka")
	case "Rasell":
		fmt.Println("Hello Rasell")
	default:
		fmt.Println("Halo", name)
	}
}
