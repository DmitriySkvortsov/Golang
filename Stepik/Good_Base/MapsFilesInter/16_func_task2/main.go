package main

import "fmt"

func task2(c chan string, s string) {
	s = fmt.Sprintf("%s ", s)
	for i := 5; i > 0; i-- {

		c <- s
	}
}
