package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines := readFile("../input.txt")
}

func checkTree(lines []string) int{

		tree :=
		for i := 0; i < 3; i++ {
			for j := 0; j < len(lines); j++ {
				tree := lines[j][i]
				if tree > 

			}
		}

	for i := len(lines[0]); i >= 0; i--{
		for j := len(lines); j > 0; j--{
			
		}
	}
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
