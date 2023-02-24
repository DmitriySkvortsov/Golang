package main

import "fmt"

func main() {
	products := map[string]int{"хлеб": 50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500}

	for name, cost := range products {
		if cost > 500 {
			fmt.Println(name, " cost more then 500")
		}

	}

	order := []string{"хлеб", "буженина", "сыр", "огурцы"}
	var sum int
	for _, o := range order {
		sum += products[o]
	}
	fmt.Println(sum)

}
