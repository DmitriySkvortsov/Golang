package main

func main() {}

func vote(x int, y int, z int) int {
	var a int
	if x == 1 {
		a += 1
	}
	if y == 1 {
		a += 1
	}
	if z == 1 {
		a += 1
	}
	if a > 1 {
		a = 1
	} else {
		a = 0
	}
	return a
}
