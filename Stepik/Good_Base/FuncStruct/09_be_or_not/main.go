package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	tx := []rune(text)

	if unicode.IsUpper(tx[0]) && string(text[len(text)-1]) == "." {
		fmt.Print("Right")
	} else {
		fmt.Print("Wrong")
	}
}
