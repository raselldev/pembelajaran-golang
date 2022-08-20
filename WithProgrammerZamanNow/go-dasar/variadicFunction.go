package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	data := sumAll(1, 2, 3)
	fmt.Println(data)
}
