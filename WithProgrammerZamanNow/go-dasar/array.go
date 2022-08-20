package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Raka"
	names[1] = "Rasell"

	fmt.Println(names[1])

	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values)
	fmt.Println(len(values))
}
