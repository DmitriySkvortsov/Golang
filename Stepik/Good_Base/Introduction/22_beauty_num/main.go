package main

import "fmt"

func main() {
	var a, b float64
	fmt.Scan(&a)
	b = a * a
	if a > 0 {

		if a > 10000 {
			fmt.Printf("%e", a)
		} else {
			fmt.Printf("%.4f", b)
		}
	} else {
		fmt.Printf("число %2.2f не подходит", a)
	}
}
