package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("perulangan ke = ", counter)
		counter++
	}

	// Test
	var hitung = 1
	for hitung = 1; hitung <= 7; {
		fmt.Println("Juara ke = ", hitung)
		hitung++
	}

}
