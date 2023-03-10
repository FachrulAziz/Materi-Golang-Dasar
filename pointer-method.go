package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	Naura := Man{"Fachrul Aziz"}
	Naura.Married()

	fmt.Println(Naura.Name)
}
