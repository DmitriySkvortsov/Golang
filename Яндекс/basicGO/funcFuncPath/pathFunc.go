package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func PrintAllFilesWithFuncFilter(path string, predicate func(string) bool) {
	var walk func(string)
	walk = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}
		for _, f := range files {
			filename := filepath.Join(path, f.Name())
			if predicate(filename) {
				fmt.Println(filename)
			}
			if f.IsDir() {
				walk(filename)
			}
		}
	}
	walk(path)
}
