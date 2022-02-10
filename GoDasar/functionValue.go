package main

import "fmt"

func helloName(name string) string {
	return "Hello" + name
}

func main() {
	sayName := helloName

	fmt.Println(sayName("Raka"))
}
