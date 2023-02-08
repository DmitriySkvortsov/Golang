package main

import "fmt"

func mult(x, y int) {
	fmt.Println(x * y)
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	mult(x, y)
}
