package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLine(filename string, proccess func(int, string)) error {
	rFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer rFile.Close()

	scanner := bufio.NewScanner(rFile)

	lineNum := 1

	for scanner.Scan() {
		proccess(lineNum, scanner.Text())
		lineNum++
	}

	return scanner.Err()
}

func main() {
	proccess := func(num int, line string) {
		fmt.Printf("%d: %s\n", num, line)
	}
	ReadLine("main.go", proccess)
}
