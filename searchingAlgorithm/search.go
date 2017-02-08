package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 22, 33, 44, 55}
	index, err := binarySearch(data, 22)
	if err != nil {
		log.Printf("Error = %v\n", err)
	}
	fmt.Printf("result index = %v\n", index)
	fmt.Printf("result = %v\n", data[index])

}

// return index of data search
func binarySearch(data []int, number int) (int, error) {
	low := 0
	high := len(data) - 1

	for high >= low {
		middle := (high + low) / 2
		if data[middle] == number {
			return middle, nil
		} else if data[middle] > number {
			high = middle
		} else {
			low = middle
		}

	}

	return -1, errors.New("Number not found!")
}
