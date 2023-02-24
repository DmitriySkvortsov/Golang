package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PrintAllFilesWithFilterClosure(path string, filter string) {
	var walk func(string)
	walk = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}
		//  проходим по списку
		for _, f := range files {
			filename := filepath.Join(path, f.Name())
			if strings.Contains(filename, filter) {
				fmt.Println(filename)
			}
			if f.IsDir() {
				walk(filename)
			}
		}
	}
	walk(path)
}
