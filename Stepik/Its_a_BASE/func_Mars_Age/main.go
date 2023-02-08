package main

import "fmt"

func main() {
	var age int
	marsAge := mars_age(age)
	fmt.Print(marsAge)
}

func mars_age(age int) int {
	age = (age * 365) / 687
	return age
}
