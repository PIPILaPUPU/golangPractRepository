package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Chdir("..")
	if err != nil {
		return
	}

	dir, err := os.Getwd()

	files, err := os.ReadDir(dir)
	fmt.Println(files)
}
