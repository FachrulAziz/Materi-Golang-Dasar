package main

import "fmt"

func namaLengkap() (namDepan, namTengah, namBelakang string) {
	namDepan = "Fachrul"
	namTengah = "Aziz"
	namBelakang = "Ganteng"
	return
}

func main() {
	a, _, c := namaLengkap() // <= tanda anderscore ( _ ) digunakan sebagai pengganti variable yang tidak dipakai
	fmt.Println(a)
	// fmt.Println(b)
	fmt.Println(c)
}
