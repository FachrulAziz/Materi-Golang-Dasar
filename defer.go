package main

import "fmt"

func logging() {
	fmt.Println("Aplikasi dijalankan")
}

func runApplikasi() {
	defer logging()
	fmt.Println("Aplikasi 1 sedang berjalan")
}

func main() {
	runApplikasi()
}
