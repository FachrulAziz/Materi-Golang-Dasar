package main

import "fmt"

func endApp() {
	fmt.Println("Aplikasi Berhenti")
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
