package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7))  // <= membulatkan float64 keatas atau kebawah, sesuai yang paling dekat
	fmt.Println(math.Round(1.3))  // <= membulatkan float64 keatas atau kebawah, sesuai yang paling dekat
	fmt.Println(math.Floor(1.7))  // <= membulatkan ke bawah
	fmt.Println(math.Ceil(1.3))   // <= membulatkan ke atas
	fmt.Println(math.Max(10, 20)) // <= mengembalikan nilai float64 paling besar
	fmt.Println(math.Min(10, 20)) // <= mengembalikan nilai float64 paling kecil
}
