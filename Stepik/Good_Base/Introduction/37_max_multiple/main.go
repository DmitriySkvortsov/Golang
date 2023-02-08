package main

import "fmt"

func main() {
	var a, b, c, res int
	fmt.Scan(&a)
	fmt.Scan(&b)
	for a <= b {
		if a%7 == 0 || a == 0 {
			c = a
			res += 1
		}
		a++
	}
	if res == 0 {
		fmt.Print("NO")
	} else {
		fmt.Print(c)
	}
}
