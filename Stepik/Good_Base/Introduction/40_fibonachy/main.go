package main

import "fmt"

func main() {
	var a, i, k, b, c int
	fmt.Scan(&a)
	k = 1
	i = 1
	for k != a {
		if k < a {
			b = k
			k = k + c
			c = b
			i++
		} else if k > a {
			i = -1
			k = a
		} else {

		}
	}
	fmt.Print(i)
}
