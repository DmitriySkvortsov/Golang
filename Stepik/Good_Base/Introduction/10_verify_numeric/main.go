package main

import "fmt"

func main() {
	var a, b, c, d int
	var w string
	fmt.Scan(&a)
	b = a / 10

	c = a % 10
	d = b / 10
	b = b % 10
	if b != c && c != d && d != b {
		w = "YES"
	} else {
		w = "NO"
	}
	fmt.Println(w)
}
