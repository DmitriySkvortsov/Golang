package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	b = 1
	for b <= a {
		fmt.Printf("%d ", b)
		b *= 2

	}
}
