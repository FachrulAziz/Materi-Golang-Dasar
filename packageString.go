package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("===== Contains =====")
	fmt.Println(strings.Contains("Fachrul Aziz", "Fachrul"))           // <= untuk mengecek apakah di string pertama ada string kedua. dan kode ini pasti menghasilkan "true"
	fmt.Println(strings.Contains("Desyana putri nurma intani", "eko")) // <= untuk mengecek apakah di string pertama ada string kedua. dan kode ini pasti menghasilkan "false"

	fmt.Println("===== Split =====")
	fmt.Println(strings.Split("Fachrul Aziz", " ")) // <= merubah data disamping mrnjadi tipe data slice

	fmt.Println("===== ToLower & ToUpper =====")
	fmt.Println(strings.ToLower("Fachrul Aziz")) // <= merubah string menjadi huruf krcil semua
	fmt.Println(strings.ToUpper("Fachrul Aziz")) // <= merubah string menjadi huruf besar semua

	fmt.Println("===== Trim =====")
	fmt.Println(strings.Trim("     Fachrul Aziz      ", " ")) // <= untuk menghapus spasi yang ada di awal dan akhir pada string pertama

	fmt.Println("===== ReplaceAll =====")
	fmt.Println(strings.ReplaceAll("aziz aziz aziz", "aziz", "Fachrul")) // <= untuk mengganti semua yang ada di string pertama dengan ketentuan string kedua dan diganti dengan string yang ketiga
}
