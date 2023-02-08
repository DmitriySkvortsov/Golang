package main

import (
	"fmt"
	"strings"
)

func main() {
	var a, c string
	fmt.Scan(&a)
	b := strings.Split(a, "")
	for i := range b {
		if i != 0 {
			c = strings.Join([]string{c, b[i]}, "*")
		} else {
			c = strings.Join([]string{c, b[i]}, "")
		}
	}
	fmt.Print(c)
}
