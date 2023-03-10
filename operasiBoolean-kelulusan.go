package main

import (
	"fmt"
)

func main() {
	// Langkah - langkah step by stepnya seperti dibawah ini
	var (
		nilaiAkhir = 90
		absensi    = 80

	// 	lulusNilaiAkhir = nilaiAkhir >= 90
	// 	lulusAbsensi    = absensi >= 80
	// 	lulus = lulusNilaiAkhir && lulusAbsensi
	)
	// fmt.Println(lulus)

	// code singkatnya seperti dibawah ini
	fmt.Println(nilaiAkhir >= 90 && absensi >= 80)
}

// Note : operator boolean ini biasa digunakan untuk kelulsan atau verivikasi boleh atau tidak
// 		  jika hasil semuanya oke, maka hasilnya oke. tetapi kalau ada hasil yang tidak oke walaupun hanya satu
//		  maka hasilnya tidak oke
// hasil dari operasi dan (&&)
// true && true = true
// true && false = false
// false && true = false
// false && false = false

// jika atau (||) maka kebalikan dari dan (&&). jika ada salah satu yang hasilnya salah, maka hasilnya tetap oke/true
// terkecuali jika semua hasilnya salah/false, maka hasil akhirnya akan salah/false.
// jika atau || kalau di artikan secara gampang, atau itu bonus yang walaupun ada 1 hasil yang salah, hasilnya tetap oke
