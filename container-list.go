package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	data.PushBack("Desyana")
	data.PushBack("Putri")
	data.PushBack("Nurma")
	data.PushBack("Intani")

	// Menampilkan data dari depan ke belakang
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	// menampilkan data dari belakang ke depan
	for elementBack := data.Back(); elementBack != nil; elementBack = elementBack.Prev() {
		fmt.Println(elementBack.Value)
	}
}
