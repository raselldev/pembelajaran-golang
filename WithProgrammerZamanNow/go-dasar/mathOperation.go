package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a + b
	fmt.Println(c)

	//augmented assignments
	var x = 10
	x += 10 // i = i + 10
	fmt.Println(x)

	//unary operation
	var i = 5
	i++
	fmt.Println(i)
}
