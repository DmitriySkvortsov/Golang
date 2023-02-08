package main

import "fmt"

func main() {
	var sum int
	fmt.Scanln(&sum)
	switch {
	case sum > 1000:
		fmt.Println("Apple")
	case sum >= 500 && sum <= 1000:
		fmt.Println("Samsung")
	case sum < 500:
		fmt.Println("Nokia с фонариком")
	}
}
