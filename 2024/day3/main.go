package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var memoryString string

func main() {
	file, err := os.Open("input.txt") // Replace "your_file.txt" with your file's name
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read the first line
	if scanner.Scan() {
		memoryString = scanner.Text()
	}

	fmt.Println("Part 1: ", part1(memoryString))
	fmt.Println("Part 2: ", part2(memoryString))
}

func part1(memory string) int {
	re := regexp.MustCompile(`(?:mul)+\((\d+(?:,\d+)*)\)`)

	//gather operations in slice
	operations := re.FindAllString(memory, -1)

	//iterate slice, extract values, perform operations
	sum := 0
	for _, v := range operations {
		first := v[strings.IndexRune(v, '(')+1 : strings.IndexRune(v, ',')]
		second := v[strings.IndexRune(v, ',')+1 : strings.IndexRune(v, ')')]

		firstInt, _ := strconv.Atoi(first)
		secondInt, _ := strconv.Atoi(second)

		sum += (firstInt * secondInt)
	}

	return sum
}

func part2(memory string) int {
	re := regexp.MustCompile(`(?:mul)+\((\d+(?:,\d+)*)\)|do\(\)|don't\(\)`)

	//gather operations in slice
	operations := re.FindAllString(memory, -1)

	//iterate slice, extract values, perform operations
	sum := 0
	enabled := true
	for _, v := range operations {
		if v == "don't()" {
			enabled = false
			continue
		}
		if v == "do()" {
			enabled = true
			continue
		}
		if enabled {
			first := v[strings.IndexRune(v, '(')+1 : strings.IndexRune(v, ',')]
			second := v[strings.IndexRune(v, ',')+1 : strings.IndexRune(v, ')')]

			firstInt, _ := strconv.Atoi(first)
			secondInt, _ := strconv.Atoi(second)

			sum += (firstInt * secondInt)
		}
	}

	return sum
}

func readFile(filepath string) []string {
	//read contents of file to lines array
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//var words []string
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Split the line into words
		//lineWords := strings.Fields(scanner.Text())
		lines = append(lines, scanner.Text())

		// Append each word to the slice
		//words = append(words, lineWords...)
	}

	return lines
}
