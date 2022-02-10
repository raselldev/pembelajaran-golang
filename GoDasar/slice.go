package main

import "fmt"

func main() {

	var month = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = month[4:7]
	//isi slice
	fmt.Println(slice1)

	//panjang slice
	fmt.Println(len(slice1))

	//kapasitas slice
	fmt.Println(cap(slice1))

	var slice2 = month[10:]
	fmt.Println(slice2)
	var slice3 = append(slice2, "Raka")
	fmt.Println(slice3)
}
