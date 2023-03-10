package main

import (
	"fmt"
)

func main() {
	var nama [3]string // tanda 3 itu adalah jumlah data yang bisa ditampung dan mulainya dimulai dari "0"
	nama[0] = "Fachrul"
	nama[1] = "Aziz"
	nama[2] = "Ganteng"

	fmt.Println(nama[2]) // saat menampilkan, kita bisa pilih ingin menampilkan data yang mana
}
