package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readFile("../input.txt")
	for _, v := range lines {

	}

	fmt.Println(lines)
}

func readFile(filepath string) []string {
	//read contents of file to lines array
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
