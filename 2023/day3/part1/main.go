package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type number struct {
	val        rune
	index      point
	partNumber bool
}

type point struct {
	x, y int
}

func main() {
	lines := readFile("../input.txt")

	numbers := make([]number, 0)
	total := 0

	for i, v := range lines {
		for j, char := range v {
			if unicode.IsDigit(char) {
				//create number object with value and position
				numb := number{
					val:   char,
					index: point{i, j},
				}

				numbers = append(numbers, numb)
				continue
			}

			val, ok := isPartNumber(numbers, lines)
			if ok {
				total += val
			}

			//reset slice
			numbers = nil
		}
	}

	fmt.Println(total)
}

func isPartNumber(numbers []number, lines [][]rune) (int, bool) {
	partnumber := false
	endValue := 0

	// Define the possible moves in the x and y directions
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for _, v := range numbers {
		for i := 0; i < 8; i++ {
			newX := v.index.x + dx[i]
			newY := v.index.y + dy[i]

			// Check if the new coordinates are within the grid boundaries
			if newX >= 0 && newY >= 0 && newX < len(lines)-1 && newY < len(lines[0])-1 {
				if lines[newX][newY] != '.' && !unicode.IsDigit(lines[newX][newY]) {
					partnumber = true
					break
				}
			}
		}
	}

	if partnumber {
		numberAsDigit := ""
		for _, v := range numbers {
			numberAsDigit += string(v.val)
		}

		endValue, _ = strconv.Atoi(numberAsDigit)
	}

	return endValue, partnumber
}

func readFile(filepath string) [][]rune {
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil
	}

	lines := strings.Split(string(content), "\n")
	result := make([][]rune, len(lines))

	for i, line := range lines {
		result[i] = []rune(line)
	}

	return result
}
