package main

import "fmt"

func main() {
	var nilaiAkhir = 80
	var absen = 80

	var lulusUjian = nilaiAkhir > 80
	var lulusAbsen = absen > 80

	var lulus = lulusUjian && lulusAbsen

	fmt.Println(lulus)
}
