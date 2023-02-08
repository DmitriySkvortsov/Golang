package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a > 0:
		fmt.Print("Число положительное")
	case a < 0:
		fmt.Print("Число отрицательное")
	default:
		fmt.Print("Ноль")
	}
}
