package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	lines := readFile("../input.txt")
	answer := 0
	for _, v := range lines {
		numList := getNum(v)

		answer += getSum(numList)
	}

	fmt.Println(answer)
}

func getNum(line string) []int {
	var result []int

	for _, v := range line {
		if unicode.IsDigit(v) {
			num, _ := strconv.Atoi(string(v))
			result = append(result, num)
		}
	}
	return result
}

func getSum(ints []int) int {
	first := ints[0] * 10
	second := ints[len(ints)-1]

	return first + second
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
