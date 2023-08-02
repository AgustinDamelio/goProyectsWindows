package main

import (
	"fmt"
)

func main() {

	testA, testB := test(22)

	fmt.Printf("Value A: %v, Value B: %v", testA, testB)
}

func test(testValue int) (a, b int) {

	if testValue > 10 {
		a := 11
		b := 12

		fmt.Printf("Value A: %v, Value B: %v inside func", a, b)
	}

	return a, b
}
