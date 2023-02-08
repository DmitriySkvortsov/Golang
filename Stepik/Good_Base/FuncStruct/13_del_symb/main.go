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

		if strings.Count(a, b[i]) > 1 {
			a = strings.Replace(a, b[i], "", -1)
		}
	}
	fmt.Print(a)

}
