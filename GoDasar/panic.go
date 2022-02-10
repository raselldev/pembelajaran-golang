package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Selesai")
	message := recover()
	fmt.Println("Terjadi error", message)
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Aplikasi Error")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
}

/**
Ini komentar
*/
