package main

import "fmt"

func main() {
	var a, sum int
	for a = 1; a <= 10; a++ {
		sum = a * a
		fmt.Println(sum)
	}
}
