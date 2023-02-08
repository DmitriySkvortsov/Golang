package main

import "fmt"

func main() {
	var a, b int
	b = 1
	for b > 0 {
		fmt.Scanln(&a)
		switch {
		case a > 100:
			b = 0
		case a >= 10 && a <= 100:
			fmt.Println(a)
		default:
			b = 1
		}
	}
}
