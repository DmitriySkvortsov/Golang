package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	b := strings.Split(a, "")
	for i := range b {
		if i%2 > 0 {
			fmt.Print(b[i])
		}
	}
}
