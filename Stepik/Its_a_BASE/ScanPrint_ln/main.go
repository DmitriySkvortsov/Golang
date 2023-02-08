package main

import "fmt"

func main() {
	var a, x, z int
	var sum, umn int

	fmt.Scanln(&a, &x, &z)

	sum = a + x + z
	umn = a * x * z

	fmt.Println(sum)
	fmt.Println(umn)
}
