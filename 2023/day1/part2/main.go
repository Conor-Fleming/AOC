package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	lines := readFile("../input.txt")
	answer := 0
	for _, v := range lines {
		digitizedString := replaceStrings(v)
		numList := getNum(digitizedString)
		answer += getSum(numList)
		fmt.Println(answer)
	}

	fmt.Println(answer)
}

func replaceStrings(line string) string {
	line = strings.ReplaceAll(line, "one", "1")
	line = strings.ReplaceAll(line, "two", "2")
	line = strings.ReplaceAll(line, "three", "3")
	line = strings.ReplaceAll(line, "four", "4")
	line = strings.ReplaceAll(line, "five", "5")
	line = strings.ReplaceAll(line, "six", "6")
	line = strings.ReplaceAll(line, "seven", "7")
	line = strings.ReplaceAll(line, "eight", "8")
	line = strings.ReplaceAll(line, "nine", "9")

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
	fmt.Println(first, second)
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
