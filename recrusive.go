package main

import (
	"fmt"
)

func factorialLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}
	return result
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	loop := factorialLoop(5)
	fmt.Println(loop)
	// fmt.Println(5 * 4 * 3 * 2 * 1) // <= code recursive diatas bila dihitung manual seperti di sebelah kiri ini

	recursive := factorialRecursive(5)
	fmt.Println(recursive)
}
