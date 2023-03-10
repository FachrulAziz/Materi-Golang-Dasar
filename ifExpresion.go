package main

import "fmt"

func main() {
	name := "Nauraaaaa"

	if name == "Fachrul" {
		fmt.Println("Hallo Ganteng")
	} else if name == "Naura" {
		fmt.Println("Hallo Naura Cantik")
	} else {
		fmt.Println("Lah elu siapa anjir ?")
	}
	// dibawah ini merupakan code batasan dari nama yang diisi.
	// jadi kalau namanya lebih dari 5 karakter, otomatis akan keluar notifikasi "Nama Terlalu Panjang"
	if karakter := len(name); karakter > 5 { // <== jangan lupa pakai len
		fmt.Println("Nama Terlalu Panjang")
	}
}
