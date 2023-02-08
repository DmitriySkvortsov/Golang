package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	max(a, b)
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
