package main

import "fmt"

func main() {
	var a int
	_, _ = fmt.Scan(&a)
	b := make([]int, a)
	_, _ = fmt.Scan(&b)
	for i := 0; i < a; i++ {
		_, _ = fmt.Scan(&b[i])
	}
	fmt.Println(b[3])
}
