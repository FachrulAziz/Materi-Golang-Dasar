package main

import "fmt"

// Tipe data Boolean hanya bisa menerima dua hasil. yaitu true/benar atau false/salah
// kode boolean di golang hanya menggunakan kata "bool" tidak seperti java "boolean"

func main() {
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	var (
		satu = 100
		dua  = 200
	)
	fmt.Println(satu > dua)
	fmt.Println(satu < dua)
	fmt.Println(satu == dua)
}
