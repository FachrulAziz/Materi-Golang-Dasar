package main

import (
	"fmt"
)

func main() {
	var name string

	name = "Fachrul Aziz"
	fmt.Println("Nama Saya " + name)

	var (
		namaDepan    = "Fachrul"
		namaBelakang = "Aziz"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)

	fmt.Println(namaDepan + " " + namaBelakang)
}
