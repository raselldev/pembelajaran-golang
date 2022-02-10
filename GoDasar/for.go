package main

import "fmt"

func main() {
	for counter := 1; counter <= 3; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	slice := []string{"raka", "rasell", "rakarasell"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}
}
