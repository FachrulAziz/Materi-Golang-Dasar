package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	result := regex.FindAllString("eko eka edo eto eyo eki", 10) // <= string eko - eki adalah data yang akan di saring dan integernya adalah jumlah data yang akan ditampilkan
	fmt.Println(result)
}
