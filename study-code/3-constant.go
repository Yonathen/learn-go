package main

import "fmt"

const TYPED_CONSTANT int = 1
const UNTYPED_CONSTANT = 2

func main() {
	const (
		A   int = 1
		PIE     = 3.14
		C       = "Area of a circle"
	)

	fmt.Println(TYPED_CONSTANT, UNTYPED_CONSTANT)
	fmt.Println(A, PIE, C)
}
