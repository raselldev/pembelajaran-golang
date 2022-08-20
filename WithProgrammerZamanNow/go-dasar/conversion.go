package main

import (
	"fmt"
)

func main() {
	var nilai32 int32 = 123456
	var nilai64 int64 = int64(nilai32)

	var nilai string = string(nilai32)

	var name = "raka"
	//saat mengambil salah satu karakter,
	//maka yang akan diambil adalah byte dari sebuah karakter
	//maka dari itu dibutuhkan konversi lagi ke sebuah string.
	var e = name[0]
	var eString = string(e)

	fmt.Println(eString)
	fmt.Println(nilai64)
	fmt.Println(nilai)
}
