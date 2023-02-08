package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c, err := divide(a, b)
	if err != nil {
		fmt.Print("ошибка")
	} else {
		fmt.Print(c)
	}
}

func divide(a int, b int) (int, error) {
	var err error
	if b == 0 {
		return b, err
	} else {
		c := a / b
		return c, err
	}
}
