package main

import "fmt"

func main() {

}

func my_print(val ...int) {
	for _, v := range val {
		fmt.Println(v)
	}
}
