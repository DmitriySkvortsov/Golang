package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a)
	b = a % 10
	c = (a / 10) % 10
	d = a / 100
	fmt.Print(d + b + c)
}
