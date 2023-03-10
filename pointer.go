package main

import "fmt"

type addres struct {
	city, province, country string
}

func main() {
	addres1 := addres{"Tangerang", "Banten", "Indonesia"}
	var addres2 = &addres1 // jika memakai tanda dan(&) maka saat di print address 1 mengikuti data adrress2
	var address3 = addres1 // jika tidak memakai tanda dan (&) maka saat di print, address 1 print data aslinya

	addres2.city = "Bandung"

	fmt.Println(addres1)
	fmt.Println(addres2)
	fmt.Println(address3)
}
