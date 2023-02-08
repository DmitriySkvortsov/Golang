package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	c := array[0]
	for j := 0; j < 5; j++ {
		if array[j] > c {
			c = array[j]
		}
	}
	fmt.Print(c)
}
