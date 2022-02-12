package main

import "fmt"

func main() {
	name := "rakaa"

	if name == "raka" {
		fmt.Println("Bener")
	} else if name == "rasell" {
		fmt.Println("Salah")
	} else {
		fmt.Println("Ya")
	}

	shortStatement()
}

func shortStatement() {
	name := "raka"
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama terlalu pendek")
	}
}
