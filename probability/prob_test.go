package main

import "testing"

func TestFactorialChair(t *testing.T) {
	type testObject struct {
		InputElement int
		InputChair   int
		Expected     int
	}

	testObjects := []testObject{
		{5, 2, 20},
		{10, 3, 720},
	}

	for _, test := range testObjects {
		actual := getCombinationSitOnChair(test.InputElement, test.InputChair)
		if actual != test.Expected {
			t.Errorf("Error actual = %v, and Expected = %v.", actual, test.Expected)
		}
	}
}

func TestGetFactorial(t *testing.T) {
	type testObject struct {
		InputNumber int
		Expected    int
	}

	testObjects := []testObject{
		{5, 120},
		{6, 720},
		{7, 5040},
	}

	for _, test := range testObjects {
		actual := getFactorial(test.InputNumber)
		if actual != test.Expected {
			t.Errorf("Error actual = %v, and Expected = %v", actual, test.Expected)
		}
	}

}

// test case 5A, should return int 5, and rune/char A
func TestGetDataBalls(t *testing.T) {
	type testObect struct {
		NumberBalls int
		TypeBalls   rune
		Input       string
	}

	testObjects := []testObect{
		{5, 'A', "5A"},
		{7, 'b', "7b"},
		{9, 'C', "9C"},
	}

	for _, test := range testObjects {

		actualNumber, actualType := getDataBalls(test.Input)

		if actualNumber != test.NumberBalls {
			t.Errorf("Error Number actual = %v, and Expected = %v", actualNumber, test.NumberBalls)
		}

		if actualType != test.TypeBalls {
			t.Errorf("Error Type actual = %v, and Expected = %v", actualType, test.TypeBalls)
		}
	}

}

// test case : 5 balls white 3 balls blue
// chance to get 2 balls white?
func TestGetChanceSameBalls(t *testing.T) {
	expected := float32(10 / 28)
	actual := getProbSameBalls(8, "2W", "5W", "3B")

	if actual != expected {
		t.Errorf("Error actual = %v, and Expected = %v", actual, expected)
	}
}
