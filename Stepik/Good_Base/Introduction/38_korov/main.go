package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a == 1 || (a > 20 && a%10 == 1):
		fmt.Printf("%d korova", a)
	case (a < 10 && (a%10 > 1 && a%10 < 5)) || (a > 20 && (a%10 > 1 && a%10 < 5)):
		fmt.Printf("%d korovy", a)
	default:
		fmt.Printf("%d korov", a)
	}
}
