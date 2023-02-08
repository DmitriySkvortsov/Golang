package main

import "fmt"

func main() {
	var ent int
	var sum int = 1
	var i int
	fmt.Scanln(&ent)
	for i = 1; i <= ent; i++ {
		sum = sum * i
	}
	fmt.Println(sum)
}
