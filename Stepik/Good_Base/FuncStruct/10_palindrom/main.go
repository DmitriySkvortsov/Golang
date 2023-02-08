package main

import (
	"fmt"
)

func main() {
	var a string
	fmt.Scanln(&a)
	b := []rune(a)
	for i := range b {
		if b[len(b)-1] == b[i] {
			fmt.Print("Палиндром")
		} else {
			fmt.Print("Нет")
		}
		break
	}
}
