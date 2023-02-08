package main

import "fmt"

func main() {
	var a, b, c, j int
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if i == 0 {
			c = b

		}
		if b == c {
			j += 1
		} else if b < c {
			c = b
			j = 1
		} else {
		}
	}
	fmt.Print(j)
}
