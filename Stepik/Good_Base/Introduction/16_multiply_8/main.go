package main

import "fmt"

func main() {
	var a, b, sum int
	fmt.Scan(&a)
	for ; a > 0; a-- {
		fmt.Scan(&b)
		if b > 10 {
			if b < 100 {
				if b%8 == 0 {

					sum += b
				}
			}
		}
	}
	fmt.Print(sum)
}
