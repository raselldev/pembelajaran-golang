package main

import "fmt"

func main() {
	const name = "Raka"

	//name = "rasell"
	//cannot assign to name (untyped string constant "Raka")
	fmt.Println(name)
}
