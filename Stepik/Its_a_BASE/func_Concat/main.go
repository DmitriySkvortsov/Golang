package main

import "fmt"

func concat(a, b string) string {
	return a + b
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	concat(a, b)
}
