package main

import "fmt"

func main() {
	//membuat alias dari type data yang sudah ada.
	type vinNum string

	var num vinNum = "1111111111"

	fmt.Println(num)
}
