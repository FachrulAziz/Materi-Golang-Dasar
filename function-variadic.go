package main

import "fmt"

// varargs adalah variable argumen
// kelebihan dari variable varadic ini adalah bisa menambahkan nilai berapapun dari satu variable
func sunAll(numbers ...int) int { // <= contohnya ini. variable number bisa diberi nilai lebih dari satu yang ada di baris ke 15
	total := 0
	for _, value := range numbers { // range numbers artinya mennjumlahkan data di barisnya
		total += value // <= total ditambahkan dengan data yang ada di barisnya (value)
	}
	return total
}
func main() {
	total := sunAll(10, 30, 35)
	fmt.Println(total)
}
