package main

import "fmt"

func main() {
	var a, b, d int
	var c string
	fmt.Scan(&a, &b, &c)
	d, c = one_or_two(a, b, c)
	fmt.Print(d, "_", c)
}

func one_or_two(a int, b int, c string) (int, string) {
	var d int
	switch c {
	case "one":
		d = a
	case "two":
		d = b
	default:
		d = 0
	}
	return d, c
}
