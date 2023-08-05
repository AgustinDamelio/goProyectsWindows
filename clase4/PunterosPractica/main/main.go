package main

import (
	"awesomeProject1/main/data"
	"fmt"
)

func main() {

	matrix, err := data.NewMatrix(3, 3, nil)

	if err != nil {
		print(err)
		return
	}

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	err = matrix.Set(data)

	if err != nil {
		print(err)
		return
	}

	matrix.Print()

	isSquare := matrix.IsSquare()

	fmt.Printf("the flag of isSquare is %v", isSquare)

	maxMatrix := matrix.Max()

	fmt.Printf("the max of matrix is %v", maxMatrix)

}
