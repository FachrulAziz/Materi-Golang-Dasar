package main

import "fmt"

type Customer struct {
	name, address string
	age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello", name, "my name is", customer.name)
}

func (a Customer) sayhuuu() {
	fmt.Println("Huuuuu from", a.name)
}

func main() {
	var fachrul Customer
	fachrul.name = "Fachrul"
	fachrul.address = "Indonesia"
	fachrul.age = 27

	fachrul.sayHi("Naura")
	fachrul.sayhuuu()
}
