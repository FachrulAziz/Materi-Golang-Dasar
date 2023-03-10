package main

import (
	"fmt"
	"os"
)

func main() {
	waw := os.Args
	fmt.Println("Argumen = ")
	fmt.Println(waw)

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("HostName = ", name)
	}
	fmt.Println("Error ", err.Error())

	username := os.Getenv("APP_PASSWORD")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}
