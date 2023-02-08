package main

import (
	"fmt"
	"strings"
)

func main() {
	var x, s, c string
	fmt.Scanln(&x, &c)
	fmt.Scanln(&s, &c)
	fmt.Println(strings.Index(x, s))
}
