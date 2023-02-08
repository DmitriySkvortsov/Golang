package main

import "fmt"

func main() {}

func minimumFromFour() int {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	switch {
	case a < b && a < c && a < d:
		return a
	case b < a && b < c && b < d:
		return b
	case c < a && c < b && c < d:
		return c
	default:
		return d
	}
}
