package main

import "fmt"

func main() {
	workArray := [10]byte{}
	var c byte
	for j := 0; j < 10; j++ {
		fmt.Scan(&workArray[j])
	}
	var i, k int
	a := [2]int{}
	i = 3
	for ; i > 0; i-- {
		fmt.Scan(&a[k], &a[k+1])
		c = workArray[a[1]]
		workArray[a[1]] = workArray[a[0]]
		workArray[a[0]] = c
	}
	for s := 0; s < 10; s++ {
		fmt.Print(workArray[s], " ")
	}
}
