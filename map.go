package main

import (
	"fmt"
)

func main() {
	person := map[string]string{
		"nama":   "Fachrul Aziz",
		"alamat": "Taman Kirana surya",
	}
	fmt.Println(person)

	delete(person, "alamat")
	fmt.Println(person)
}
