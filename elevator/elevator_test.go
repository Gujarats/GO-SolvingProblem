package main

import (
	"fmt"
	"testing"
)

func TestSoulution(t *testing.T) {
	testObjects := []struct {
		Persons     []int
		Lift        []int
		MaxLift     int
		TotalPerson int
		MaxWeight   int
	}{
		// 0
		{
			Persons:     []int{60, 50, 20, 40, 60},
			Lift:        []int{1, 2, 3, 1, 1},
			MaxLift:     10,
			TotalPerson: 2,
			MaxWeight:   200,
		},

		// 1
		{
			Persons:     []int{60, 50, 20, 40, 60},
			Lift:        []int{1, 1, 3, 3, 0},
			MaxLift:     10,
			TotalPerson: 2,
			MaxWeight:   200,
		},

		// 2
		{
			Persons:     []int{60, 50, 20, 40, 60},
			Lift:        []int{2, 2, 2, 3, 3},
			MaxLift:     10,
			TotalPerson: 1,
			MaxWeight:   200,
		},

		// 3
		{
			Persons:     []int{60, 50, 20, 40, 60},
			Lift:        []int{2, 2, 2, 2, 2},
			MaxLift:     10,
			TotalPerson: 5,
			MaxWeight:   2000,
		},

		// 4
		{
			Persons:     []int{60, 50, 20, 40, 60},
			Lift:        []int{2, 2, 2, 2, 2},
			MaxLift:     1,
			TotalPerson: 2,
			MaxWeight:   200,
		},
	}

	for index, testObject := range testObjects {
		fmt.Println("running test index = ", index)
		actual := solutions(testObject.Persons, testObject.Lift, testObject.MaxLift, testObject.TotalPerson, testObject.MaxWeight)
		fmt.Println("result = ", actual)
	}
}
