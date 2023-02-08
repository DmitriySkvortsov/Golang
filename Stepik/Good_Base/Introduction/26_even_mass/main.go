package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Scanln(&a, &b)
	if a >= 2 {

		for j := 0; j < a; j++ {
			if j%2 == 0 {
				fmt.Scan(&b, &c)
				fmt.Print(b, " ")
			}
		}
	} else {
		fmt.Scan(&b)
		fmt.Print(b)
	}
}
