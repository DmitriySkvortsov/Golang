package main

import (
	"fmt"
)

func main() {
	var a string
	var d rune
	fmt.Scan(&a)
	r := []rune(a)
	for i := range r {
		if r[i] > d {
			d = r[i]
		}
	}
	g := string(d)

	fmt.Print(g)
}
