//struct adalah kumpulan dari beberata tipe data
//yang disimpan dalam 1 kesatuan

package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func (x Customer) sayHi() {
	fmt.Println("Hai from", x.name)
}

func main() {
	//menggunakan struct literals
	rasell := Customer{
		name:    "rasell",
		address: "jkt",
		age:     20,
	}

	fmt.Println(rasell)

	//menggunakan struct biasa
	var raka Customer
	raka.name = "raka"
	raka.address = "jakarta"
	raka.age = 30

	fmt.Print(raka)
	raka.sayHi()
}
