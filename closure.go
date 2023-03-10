package main

import (
	"fmt"
)

func main() {
	name := "Fachrul Aziz"
	counter := 0
	/* closure yang dimaksud disini adalah, function yang dibawah ini bisa mengambil data dari function yang ada di atas,
	tetapi function yang ada di atas tidak bisa akses data yang ada di function dibaeah ini */
	increment := func() {
		name1 := "Naura"
		fmt.Println(name)
		counter++
		fmt.Println(name1)
	}

	increment()
	increment()
}
