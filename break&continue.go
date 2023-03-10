package main

import "fmt"

func main() {
	// break merupakan kata kunci untuk menghentikan perhitungan
	for i := 1; i <= 10; i++ {
		fmt.Println("Perhitungan ke =", i)
		if i == 7 { // <= jadi, ketika perhitungan sudah sampai 7, perhitungan di hentikan
			break
		}
	}

	// continue merupakan kata kunci untuk menghentikan yang dimaksud tetapi terus menghitung
	// di baris ke 17, merupakan module yang maksudnya adalah jika 2 dibagi 0 maka yang ditampilkan hasil ganjil dan jika 2 dibagi 1, maka yang ditampilkan hasil genap
	for a := 1; a <= 10; a++ {
		if a%2 == 0 { // <= disamping merupakan perintah jika perhitungan menghitung angka genap, maka tidak di hitung dan malah melompat ke angka selanjutnya
			continue
		}
		fmt.Println("Ke = ", a)

	}
}
