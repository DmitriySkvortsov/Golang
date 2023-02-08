package main

import "fmt"

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

func sumInt(a ...int) (b int, c int) {
	b = len(a)
	for i := 0; i < b; i++ {
		c = c + a[i]
	}
	return b, c
}
