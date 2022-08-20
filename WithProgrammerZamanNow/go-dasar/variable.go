package main

import "fmt"

func main() {
	var name string

	name = "Raka"
	fmt.Println(name)

	var country = "Indonesia"
	fmt.Println(country)

	//Deklarasi awal saja jika ingin menggunakan :=
	//Selanjutnya menggunakan =  biasa
	age := 24
	fmt.Println(age)

	var (
		firstName = "Raka"
		lastName  = "Rasell"
	)

	fmt.Println(firstName + " " + lastName)
}
