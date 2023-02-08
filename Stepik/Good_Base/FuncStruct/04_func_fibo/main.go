package main

func main() {}

func fibonacci(n int) int {
	var i, k, b, c int

	k = 1

	for i < n-1 {
		b = k
		k = k + c
		c = b
		i++
	}
	return k
}
