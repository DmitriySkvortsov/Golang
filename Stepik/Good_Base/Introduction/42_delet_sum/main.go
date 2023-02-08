package main

import "fmt"

func main() {
	var a, b, i int
	var arr [100]int
	fmt.Scan(&a)
	fmt.Scan(&b)
	for a > 0 {
		if a%10 != b {
			if a > 10 {
				arr[i] = a % 10

				a /= 10
			} else {
				arr[i] = a
				a = 0
			}
			i += 1
		} else {
			a /= 10
		}
	}

	for j := i - 1; j >= 0; j-- {
		fmt.Print(arr[j])
	}
}
