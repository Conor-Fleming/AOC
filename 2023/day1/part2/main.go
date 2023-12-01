package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var digitStrings = map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}

func main() {
	lines := readFile("../input.txt")
	answer := 0
	for _, v := range lines {
		digitizedString := replaceStrings(v)
		fmt.Println(v)
		numList := getNum(digitizedString)
		fmt.Println(numList)
		answer += getSum(numList)
		fmt.Println(answer)
	}

	fmt.Println(answer)
}

func replaceStrings(line string) string {
	digString := ""
	for _, v := range line {
		if !unicode.IsLetter(v) {
			digString = ""
			continue
		}
		digString += string(v)

		if v, ok := digitStrings[digString]; ok {
			line = strings.Replace(line, digString, v, 1)
		}
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
