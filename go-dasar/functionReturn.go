package main

import "fmt"

func getReturn(name string) string {
	return "hello" + name
}

func main() {
	result := getReturn("raka")

	fmt.Println(result)
}
