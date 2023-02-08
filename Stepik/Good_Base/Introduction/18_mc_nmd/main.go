package main

import "fmt"

func main() {
	var a, b, c, i int16
	fmt.Scan(&a, &b, &c)

	i = 1
	for ; i <= a; i++ {
		if i%b == 0 && i%c != 0 {
			fmt.Print(i)
			i = a
		}
	}
}
