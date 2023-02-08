package main

import "fmt"

func main() {
	var a, b, sum int
	fmt.Scan(&a, &b)
	sum = double_m(a, b)
	fmt.Print(sum)
}

func double_m(a int, b int) int {
	var sum int
	var d int
	var e int
	if a > b {
		e = a
		d = b
	} else {
		e = b
		d = a
	}

	for ; e >= d; e-- {
		if e == d {
			sum += (e * e)
		} else {
			sum = sum + (d * d) + (e * e)
			d += 1
		}
	}
	return sum
}
