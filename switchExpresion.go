package main

import "fmt"

func main() {
	nama := "Fachrul"

	switch nama {
	case "Fachrul":
		fmt.Println("Hallo Fachrul Aziz")
	case "Naura":
		fmt.Println("Hallo Naura Almeera Fachrul anak ayah Fachrul Aziz")
	default:
		fmt.Println("Hallo Bayi tirexs")
	}

	// switch karakter := len(nama); karakter > 5 {
	// case true:
	// 	fmt.Println("Nama Terlalu Panjang")
	// }

	karakter := len(nama) // len merupakan code untuk menghitung jumlah karakter di string
	switch {
	case karakter > 10:
		fmt.Println("Nama Terlalu Panjang")
	case karakter > 7:
		fmt.Println("Nama Lumayan Panjang")
	case karakter > 3:
		fmt.Println("Nama Oke")
	}
}
