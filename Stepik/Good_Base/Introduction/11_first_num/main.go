package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	for a > 10 {
		a /= 10
	}
	if a == 10 {
		a = 1
	}
	fmt.Print(a)
}
