package main

import "fmt"

func main() {}

func test(x1 *int, x2 *int) {
	*x1 = *x1 * *x2
	fmt.Print(*x1)
}
