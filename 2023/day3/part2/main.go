package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type point struct {
	x, y int
}

func main() {
	lines := readFile("../input.txt")

	total := 0

	for i, v := range lines {
		for j, char := range v {
			if char == '*' {
				point := point{i, j}

				val, ok := isGear(point, lines)
				if ok {
					total += val
				}
			}
		}
	}

	fmt.Println(total)
}

func isGear(star point, lines [][]rune) (int, bool) {
	numbers := make(map[int]bool)

	// Define the possible moves in the x and y directions
	dx := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dy := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	for i := 0; i < 8; i++ {
		newX := star.x + dx[i]
		newY := star.y + dy[i]

		// Check if the new coordinates are within the grid boundaries
		if newX >= 0 && newY >= 0 && newX < len(lines)-1 && newY < len(lines[0])-1 {
			start := point{newX, newY}
			if unicode.IsDigit(lines[newX][newY]) {
				number := getNumber(start, lines)
				fmt.Println(string(number))
				if _, ok := numbers[number]; !ok {
					numbers[number] = true
				}
			}
		}
	}

	if len(numbers) != 2 {
		return 0, false
	}

	ratio := 1
	for k := range numbers {
		ratio *= k
	}

	return ratio, true
}

func getNumber(startPos point, grid [][]rune) int {
	number := ""
	dy := []int{-2, -1, 1, 2}

	for i := 0; i < 4; i++ {
		indexCheck := startPos.y + dy[i]
		if indexCheck >= 0 && indexCheck < len(grid[startPos.x]) {
			val := grid[startPos.x][indexCheck]
			if unicode.IsDigit(val) {
				number += string(val)
			}
		}
	}

	resultNumber, _ := strconv.Atoi(number)

	return resultNumber
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
