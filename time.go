package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())

	fmt.Println("=========================")
	formatTanggal := time.Date(2023, 03, 9, 10, 35, 57, 20, time.UTC)
	fmt.Println(formatTanggal)

	fmt.Println("=========================")
	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2023-03-09")
	fmt.Println(parse)
}
