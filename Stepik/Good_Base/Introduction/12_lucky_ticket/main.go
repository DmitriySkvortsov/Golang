package main

import "fmt"

func main() {
	var a, b, suma, sumb, d int
	fmt.Scan(&a)
	b = a
	a /= 1000
	for a > 10 {
		d = a % 10
		suma += d
		a /= 10
	}
	suma += a
	b = b % 1000
	for b > 10 {
		d = b % 10
		sumb += d
		b /= 10
	}
	sumb += b

	if suma == sumb {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
