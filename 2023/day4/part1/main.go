package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readFile("../input.txt")
	total := 0
	for _, v := range lines {
		numbers := strings.Split(v, ":")
		values := strings.Fields(numbers[1])

		total += findMatches(values)
	}

	fmt.Println(total)
}

func findMatches(numberList []string) int {
	matches := make(map[string]bool)
	totalMatches := 0

	for _, v := range numberList {
		if v == "|" {
			continue
		}

		if _, ok := matches[v]; ok {
			totalMatches++
			continue
		}

		matches[string(v)] = true
	}

	if totalMatches == 0 {
		return 0
	}

	total := 1
	for i := 1; i < totalMatches; i++ {
		total *= 2
	}

	return total
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

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
