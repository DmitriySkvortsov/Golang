package main

import "fmt"

func main() {

	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	var nr string
	fmt.Scanln(&nr)
	results = append(results, nr)
	var points int
	for _, res := range results {
		switch res {
		case "w":
			points += 3
		case "d":
			points += 1
		default:
			points += 0
		}
	}
	fmt.Println(points)
}
