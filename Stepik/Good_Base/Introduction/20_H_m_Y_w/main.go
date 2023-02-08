package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c)
	for a < c {
		a = a + (a * b / 100)
		d += 1
	}
	fmt.Print(d)
}
