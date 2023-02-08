package main

import "fmt"

func main() {
	var a, b float32
	fmt.Scan(&a, &b)
	var c float32
	c = (a + b) / 2
	fmt.Print(c)
}
