package main

import (
	"errors"
	"fmt"
)

func main() {
	// create matrix a and b
	type Matrix [][]int
	a := Matrix{
		{2, 3, 4},
		{2, 3, 4},
		{2, 3, 4},
		{2, 3, 4},
	}

	b := Matrix{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	result, _ := multiplication(a[:][:], b[:][:])

	fmt.Printf("resutl = %+v\n", result)
}

// Rules multiplications is a rows * b column
// requirement aColumn == bRow
func multiplication(a [][]int, b [][]int) ([][]int, error) {
	aRow := len(a)
	aCol := len(a[0])
	bRow := len(b)
	bCol := len(b[0])
	fmt.Printf("a = %+v\n", a)
	fmt.Printf("b = %+v\n", b)

	//checkin the compatibility matrix to multiply
	if aCol != bRow {
		newError := errors.New("Matrix is not compatible")
		return nil, newError
	}

	//var result [aRow][bCol]int
	result := make([][]int, aRow)
	for i := range result {
		result[i] = make([]int, bCol)
	}

	for row := 0; row < aRow; row++ {
		for column := 0; column < aCol; column++ {
			for colb := 0; colb < bCol; colb++ {
				result[row][column] += a[row][column] * b[column][colb]
			}
		}
	}

	return result, nil
}
