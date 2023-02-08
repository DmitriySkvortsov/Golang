package main

import "fmt"

func main() {}

func test(x1 *int, x2 *int) {
	var c int
	c = *x1
	*x1 = *x2
	*x2 = c
	fmt.Print(*x1, *x2)
}
