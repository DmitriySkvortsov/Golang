package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b, c := calc(a)
	fmt.Print(&b, &c)
}

func calc(a int) (int, int) {
	return a * 2, a * a
}
