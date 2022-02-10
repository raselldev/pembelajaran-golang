package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Raka",
		"address": "JKT",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
