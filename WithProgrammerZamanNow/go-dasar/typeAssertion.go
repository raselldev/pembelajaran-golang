package main

import "fmt"

func random() interface{} {
	return "ok"
}

func main() {
	result := random()
	resultString := result.(string)
	fmt.Println(resultString)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknow Type")
	}
}
