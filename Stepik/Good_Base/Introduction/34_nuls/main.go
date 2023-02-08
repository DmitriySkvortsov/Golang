package main

import "fmt"

func main() {
	var a, b, j int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b == 0 {
			j += 1
		}
	}
	fmt.Print(j)
}
