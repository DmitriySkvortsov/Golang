package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	isEven(a)
}

func isEven(a int) bool {
	var c int
	c = a % 2
	if c == 0 {
		return true
	} else {
		return false
	}
}
