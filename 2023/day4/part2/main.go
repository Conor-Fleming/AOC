package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readFile("../input.txt")
	cards := make(map[int]int, len(lines))

	for i := range lines {
		cards[i] = 1
	}

	for i, val := range lines {
		for j := 0; j < cards[i]; j++ {
			matches := 0
			numbers := strings.Split(val, ":")
			values := strings.Fields(numbers[1])

			matches += findMatches(values)
			for y := i + 1; y <= i+matches; y++ {
				cards[y] += 1
			}
		}
	}

	sum := 0
	for _, v := range cards {
		sum += v
	}
	fmt.Println(sum)
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

	return totalMatches
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
