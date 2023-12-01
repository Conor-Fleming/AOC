package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitStrings = map[string]string{"one": "o1e", "two": "t2o", "three": "t3e", "four": "f4r", "five": "f5e", "six": "s6x", "seven": "s7n", "eight": "e8t", "nine": "n9e"}

func main() {
	lines := readFile("../input.txt")
	answer := 0

	for _, v := range lines {
		digitizedString := replaceStrings(v)
		numList := getNum(digitizedString)
		answer += getSum(numList)
	}

	fmt.Println(answer)
}

func replaceStrings(line string) string {
	for i := range digitStrings {
		line = strings.Replace(line, i, digitStrings[i], -1)
	}

	return line
}

func getNum(line string) []int {
	var result []int
	lineWithDigits := replaceStrings(line)
	for _, v := range lineWithDigits {
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
