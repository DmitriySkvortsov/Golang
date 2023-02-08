package main

import (
	"fmt"
	"time"
)

func main() {
	var a float32
	var b, c int
	b = 10
	m := make(map[float32]int)
	for b > 0 {
		fmt.Scan(&a)
		if m[a] > 0 {
			fmt.Print(m[a])
		} else {
			c = int(a)
			m[a] = work(c)
			fmt.Print(m[a])
		}
		fmt.Print(" ")
		b -= 1
	}
}

func work(n int) int {
	if n > 3 {
		time.Sleep(time.Millisecond * 500)
		return n + 1
	} else {
		time.Sleep(time.Millisecond * 500)
		return n - 1
	}
}
