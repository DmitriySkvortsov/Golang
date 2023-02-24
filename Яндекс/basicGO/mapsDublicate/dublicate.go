package main

import "fmt"

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}

	ok := RemoveDuplicates(input)

	fmt.Println(ok)
}

func RemoveDuplicates(input []string) []string {

	allKeys := make(map[string]bool)
	ok := []string{}
	for _, item := range input {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			ok = append(ok, item)
		}
	}
	return ok
}
