package main

import "fmt"

func main() {
	var a, c int
	fmt.Scan(&a)
	a = a / 10
	c = a % 10
	fmt.Print(c)

}
