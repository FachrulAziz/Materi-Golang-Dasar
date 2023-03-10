package main

import "fmt"

type Filter func(string) string

func sayHelloFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello " + nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else if name == "monyet" {
		return "..."
	} else if name == "babi" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloFilter("Fachrul Aziz", spamFilter)
	sayHelloFilter("monyet", spamFilter)
}
