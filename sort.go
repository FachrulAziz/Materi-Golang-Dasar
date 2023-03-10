package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	Users := []User{
		{"Fachrul Aziz", 27},
		{"Ikbal Maulana", 31},
		{"Naura ameera", 2},
		{"Daffara ALendra", 1},
	}
	fmt.Println(Users)
	sort.Sort(UserSlice(Users))
	fmt.Println(Users)
}
