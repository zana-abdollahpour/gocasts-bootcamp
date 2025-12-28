package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	var file string
	if len(args) <= 1 {
		fmt.Print("no file-name provided, please enter target filename: ")
		fmt.Scan(&file)
	} else {
		file = args[1]
	}

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		if fileInfo, err := os.Stat(fullPath); err == nil {
			mode := fileInfo.Mode()
			if mode.IsRegular() && mode&0111 != 0 {
				fmt.Println(fullPath)
				return
			}
		}
	}
	fmt.Println("NOT FOUND!")
}
