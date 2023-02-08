package main

import "fmt"

func main() {
	var b, c, d int

	for fmt.Scan(&b); b != 0; fmt.Scan(&b) {

		switch {
		case b > c:
			c = b
			d = 1
		case b == c:
			d += 1
		}
	}
	fmt.Print(d)
}
