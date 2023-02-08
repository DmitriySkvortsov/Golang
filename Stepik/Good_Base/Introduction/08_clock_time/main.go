package main

import "fmt"

func main() {
	var a, c, o, m int
	fmt.Scan(&a)
	c = a / 30
	o = a % 30
	m = o * 2
	fmt.Println("It is", c, "hours", m, "minutes.")
}
