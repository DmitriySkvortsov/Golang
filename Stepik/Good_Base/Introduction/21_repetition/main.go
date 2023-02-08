package main

import (
	"fmt"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	c := []rune(a)
	d := []rune(b)
	for i := range c {
		for j := range d {
			if c[i] == d[j] {
				fmt.Printf("%s ", string(c[i]))
			}
		}
	}
}
