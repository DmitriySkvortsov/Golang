package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a)
	b = a / 3600
	c = (a % 3600) / 60
	fmt.Printf("It is %d hours %d minutes.", b, c)
}
