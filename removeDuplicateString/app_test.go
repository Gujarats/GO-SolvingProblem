package main

import (
	"testing"
)

func TestDuplicateString(t *testing.T) {
	type testObject struct {
		Input    string
		Expected string
	}

	testObjects := []testObject{
		{"asdfasdf", "asdf"},
		{"maahhnn", "mahn"},
		{"aabbcc", "abc"},
	}

	for _, test := range testObjects {
		actual := removeDulicateString(test.Input)
		if test.Expected != actual {
			t.Errorf("Error actual = %v, Expected = %v", actual, test.Expected)

		}
	}
}

func TestRemoeSpecifiArray(t *testing.T) {
	type testObject struct {
		InputData []rune
		Index     int
		Expected  []rune
	}

	testObjects := []testObject{
		{[]rune{'a', 'b', 'c'}, 2, []rune{'a', 'b'}},
		{[]rune{'a', 'b', 'c'}, 1, []rune{'a', 'c'}},
		{[]rune{'a', 'b', 'c'}, 0, []rune{'b', 'c'}},
	}

	for _, test := range testObjects {
		actual := removeSpecifiArray(test.InputData, test.Index)
		if !testEq(actual, test.Expected) {
			t.Errorf("Error actual = %v, Expected = %v", actual, test.Expected)

		}
	}
}

// checking if slice a is equal to slice b.
func testEq(a, b []rune) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
