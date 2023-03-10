package main

import "fmt"

func main() {

	type noKTP string
	type married bool

	var nomorKTPEko noKTP = "3241243324"
	var marriedStatus married = true
	fmt.Println(nomorKTPEko)
	fmt.Println(marriedStatus)
	var nomorKTPNaura noKTP = "43232423432"
	fmt.Println(nomorKTPNaura)
}
