package main

import (
	"fmt"
	"strings"
)

func main() {
	a := 'q'
	b := "aq"
	c := []rune(b)
	d := "q"

	fmt.Printf("a = %#v\n", a)
	fmt.Printf("b = %#v\n", b)
	fmt.Printf("c = %#v\n", c)
	fmt.Printf("d = %#v\n", d)

	if a == (c[1]) {
		fmt.Printf("value %+v is equal %+v\n", a, c[1])
	} else {
		fmt.Printf("value %+v is NOT %+v\n", a, c[1])
	}

	if d == strings.ToLower(string(c[1])) {
		fmt.Printf("value %+v is equal %+v\n", d, strings.ToLower(string(c[1])))
	} else {
		fmt.Printf("value %+v is NOT %+v\n", d, string(c[1]))
	}

}
