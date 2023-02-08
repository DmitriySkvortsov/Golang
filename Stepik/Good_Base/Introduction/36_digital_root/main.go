package main

import "fmt"

func main() {
	var a, c int
	fmt.Scan(&a)
	for a%10 > 0 {
		c += a % 10
		if a > 10 {
			a /= 10
		} else {
			if c >= 10 {
				a = c
				c = 0
			} else {
				a = 0
			}
		}
	}
	fmt.Print(c)
}
