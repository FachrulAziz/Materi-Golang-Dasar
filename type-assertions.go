package main

import "fmt"

/*Tipe Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan.
Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong */
func random() interface{} {
	return "Fachrul Aziz"
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)
	switch value := result.(type) {
	case string:
		fmt.Println("value", value, "Is String")
	case int:
		fmt.Println("value", value, "Is Integer")
	default:
		fmt.Println("Unknown type")
	}
}
