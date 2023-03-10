package main

import "fmt"

func main() {
	var bulan = []string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"Nobember",
		"Desember",
	}
	var slice1 = bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// bulan[5] = "Satu"
	// fmt.Println(slice1)
}
