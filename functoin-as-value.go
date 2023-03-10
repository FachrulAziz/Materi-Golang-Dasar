package main

import "fmt"

func getGoodBye(name string) string {
	return "good bye " + name
}

func main() {
	kwkw := getGoodBye
	//	result := kwkw("Fachrul Aziz")
	//	fmt.Println(result)
	fmt.Println(kwkw("Fachrul Aziz"))
}
