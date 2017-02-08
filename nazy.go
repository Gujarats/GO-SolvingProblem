package main

import (
	"fmt"
)

func main() {
	solutions(5)
}

func print2DArray(nazy [][]string) {
	for i := 0; i < len(nazy); i++ {
		for j := 0; j < len(nazy[0]); j++ {
			fmt.Print(nazy[i][j], " ")
		}
		fmt.Println()
	}
}

func solutions(input int) {
	//initiate 2d array
	var nazy = make([][]string, input)
	for i := range nazy {
		nazy[i] = make([]string, input)
	}

	nazy = assignEmptySpace(nazy)

	middle := len(nazy) / 2
	fmt.Println("Middle = ", middle)
	fmt.Println("array lenght = ", len(nazy))
	fmt.Println("array lenght2 = ", len(nazy[0]))

	//first line
	for i := 0; i <= middle; i++ {
		nazy[0][i] = "*"
	}

	//middle line
	for i := 0; i < input; i++ {
		nazy[i][middle] = "*"
	}

	//finish line
	for i := middle; i < input; i++ {
		nazy[len(nazy)-1][i] = "*"
	}

	// second

	//first line
	for i := len(nazy) - 1; i >= middle; i-- {
		nazy[i][0] = "*"
	}

	//middle line
	for i := 0; i < len(nazy); i++ {
		nazy[middle][i] = "*"
	}

	// finish line
	for i := middle; i >= 0; i-- {
		nazy[i][len(nazy)-1] = "*"
	}

	print2DArray(nazy)
}

func assignEmptySpace(nazy [][]string) [][]string {
	for i := 0; i < len(nazy); i++ {
		for j := 0; j < len(nazy[0]); j++ {
			nazy[i][j] = " "
		}
	}

	return nazy
}
