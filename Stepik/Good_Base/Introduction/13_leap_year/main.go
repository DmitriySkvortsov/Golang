package main

import "fmt"

func main() {
	var a int
	var b bool
	fmt.Scan(&a)
	b = a%4 == 0 && a%100 > 0
	switch {
	case b || a%400 == 0:
		fmt.Print("YES")
	default:
		fmt.Print("NO")
	}
}
