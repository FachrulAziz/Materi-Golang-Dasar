package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello ", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var Fachrul Person
	Fachrul.Name = "Fachrul Aziz"

	sayHello(Fachrul)

	cat := Animal{
		Name: "Push",
	}
	sayHello(cat)
}
