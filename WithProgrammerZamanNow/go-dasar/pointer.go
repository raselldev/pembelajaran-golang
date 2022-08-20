package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func changeCountry(address Address) {
	address.Negara = "Amerika"
}

func main() {
	address1 := Address{"Jakarta", "DKI", "Indonesia"}
	//pass to value
	address2 := address1

	//pass to reference
	//address2 adalah pointer dari address1
	//address2 := &address1

	address2.Kota = "Bandung"

	fmt.Println(address2)

	//pointer kosong menggunakan new
	address4 := new(Address)
	fmt.Println(address4)
}
