package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	value, _ := strconv.Atoi("20000")
	fmt.Println(value)
	fmt.Println("fachrul", value)
	value1 := strconv.Itoa(2000)
	fmt.Println(value1)
}
