package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scan(&a)
	b := []rune(a)
	for i := range b {

		d := (b[i] - 48) * (b[i] - 48)
		fmt.Print(d)

	}

}
