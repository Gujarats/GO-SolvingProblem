package main

import (
	"testing"
)

func TestReplaceChat(t *testing.T) {
	type testObject struct {
		Input    string
		Expected string
	}
	var testObjects = []testObject{
		{"tet", ""},
		{"telot", "lo"},
		{"telolet", "lole"},
		{"telolelot", "lolelo"},
		{"telolelolet", "lolelole"},
		{"telolelolelot", "lolelolelo"},
	}

	for _, test := range testObjects {
		actual := ReplaceChar(test.Input)
		if actual != test.Expected {
			t.Errorf("Test Failed expected : %s, actual : %s\n", test.Expected, actual)
		}
	}

}

func TestIsLo(t *testing.T) {
	type testObject struct {
		Input    string
		Expected bool
	}
	var testObjects = []testObject{
		{"lo", true},
		{"", false},
		{"le", false},
		{"lolelo", true},
		{"lolelole", false},
	}

	for _, test := range testObjects {
		actual := isLastLo(test.Input)
		if actual != test.Expected {
			t.Errorf("Test Failed expected : %v, actulal : %v\n", actual, test.Expected)
		}
	}
}

func TestAddLole(t *testing.T) {
	type testObject struct {
		Input       string
		Inputbool   bool
		InputNumber int
		Expected    string
	}

	var testObjects = []testObject{
		{"lole", false, 1, "lolelo"},
		{"lole", false, 2, "lolelole"},
		{"lole", false, 3, "lolelolelo"},
		{"le", false, 1, "lelo"},
		{"le", false, 2, "lelole"},
		{"", false, 1, "lo"},
		{"", false, 2, "lole"},
		{"", false, 3, "lolelo"},
	}

	for _, test := range testObjects {
		addLole(&test.Input, test.Inputbool, test.InputNumber)
		if test.Input != test.Expected {
			t.Errorf("Test Failed expected : %v, actulal : %v\n", test.Expected, test.Input)
		}
	}
}

func TestTelolet(t *testing.T) {
	type testObject struct {
		Input       string
		InputNumber int
		Expected    string
	}

	var testObjects = []testObject{
		{"", 2, "telolet"},
		{"", 3, "telolelot"},
		{"", 4, "telolelolet"},
		{"telolet", 1, "telolelot"},
		{"telolet", 2, "telolelolet"},
		{"telolet", 3, "telolelolelot"},
		{"telolet", 4, "telolelolelolet"},
		{"telolelot", 1, "telolelolet"},
		{"telolelot", 2, "telolelolelot"},
		{"telolelot", 3, "telolelolelolet"},
		{"tet", 1, "telot"},
		{"tet", 2, "telolet"},
		{"tet", 3, "telolelot"},
		{"tet", 4, "telolelolet"},
		{"tet", 5, "telolelolelot"},
		{"tet", 6, "telolelolelolet"},
		{"tet", 7, "telolelolelolelot"},
	}

	for _, test := range testObjects {
		telolet(&test.Input, test.InputNumber)
		if test.Input != test.Expected {
			t.Errorf("Test Failed expected : %v, actulal : %v\n", test.Expected, test.Input)
		}
	}
}
