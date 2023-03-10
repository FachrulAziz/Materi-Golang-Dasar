package main

import (
	"Programmer-Jaman-Now/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
