package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
   Determine the ASCII code for the current character of the string.
   Increase the current value by the ASCII code you just determined.
   Set the current value to itself multiplied by 17.
   Set the current value to the remainder of dividing itself by 256.x
*/

func main() {
	steps := readFile("../input.txt")

	total := 0

	for _, v := range steps {
		total += hashIt(v)
	}

	fmt.Println(total)
}

func hashIt(step string) int {
	curVal := 0
	for _, v := range step {
		curVal += int(v)
		curVal *= 17
		curVal %= 256
	}

	return curVal
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
		lines = append(lines, strings.Split(scanner.Text(), ",")...)
	}

	return lines
}
