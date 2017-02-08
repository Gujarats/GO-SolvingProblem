package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	a := "telolet"
	telolet(&a, 10)
	fmt.Println("result : ", a)
}

// here we are going to replace te with empty string,
// and t with empty string
func ReplaceChar(telolet string) string {
	replacer := strings.NewReplacer("te", "", "t", "")
	return replacer.Replace(telolet)
}

// Get the las char of input params
// compare it with o char
func isLastLo(lole string) bool {
	arrayLole := []rune(lole)

	if len(arrayLole) == 0 {
		return false
	}

	// get the last char conver it to string and lower case
	if strings.ToLower(string(arrayLole[utf8.RuneCountInString(lole)-1])) != "o" {
		return false
	}

	return true
}

func addLole(lole *string, isLo bool, number int) {
	for i := 0; i < number; i++ {
		if isLo {
			isLo = false
			*lole += "le"
		} else {
			isLo = true
			*lole += "lo"
		}
	}
}

func telolet(telolet *string, number int) {
	prefix := "te"
	last := "t"

	lole := ReplaceChar(*telolet)

	lo := isLastLo(lole)

	addLole(&lole, lo, number)

	*telolet = prefix + lole + last

}
