package main

func task(c chan int, a int) {
	a += 1
	c <- a
}
