package main

import "fmt"

func main() {
	var gender int
	switch gender {
	case 1:
		fmt.Println("Мужчина")
	case 2:
		fmt.Println("Женщина")
	default:
		fmt.Println("Еще не определился")
	}
}
