package main

import "fmt"

func Generate(seed int) func() {
	return func() {
		fmt.Println(seed)
		seed += 2
	}

}

func main() {
	iterator := Generate(0)
	iterator()
	iterator()
	iterator()
	iterator()
	iterator()
}
