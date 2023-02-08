package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	out := adding(a, b)
	fmt.Print(out)
}

func adding(a string, c string) int64 {
	var one, two string
	b := []rune(a)
	d := []rune(c)

	for i := range b {

		if unicode.IsDigit(b[i]) {
			h := string(b[i])
			one = strings.Join([]string{one, h}, "")
		}
	}
	for j := range d {
		if unicode.IsDigit(d[j]) {
			k := string(d[j])
			two = strings.Join([]string{two, k}, "")
		}
	}

	ones, _ := strconv.ParseInt(one, 10, 64)
	twos, _ := strconv.ParseInt(two, 10, 64)
	e := ones + twos

	return e
}
