package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh NOL")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
	hasil, err := pembagi(200, 0)
	if err == nil {
		fmt.Println("Hasil ", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}
}
