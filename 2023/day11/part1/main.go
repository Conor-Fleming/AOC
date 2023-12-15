package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid := readFile("../input.txt")
	grid = expandUniv(grid)

	for _, v := range grid {
		fmt.Println(string(v))
	}

}

func expandUniv(grid [][]rune) [][]rune {
	for i, v := range grid {
		if !strings.Contains(string(v), "#") {
			firstPart := append([][]rune{}, grid[:i+1]...)
			secondPart := append([][]rune{}, grid[i+1:]...)
			firstPart = append(firstPart, v)
			grid = append(firstPart, secondPart...)
		}
	}

	for i := range grid {
		for j, val := range grid {
			if grid[j][i] == '#' {
				fmt.Println("here")
				for i := range grid {
					firstPart := append([]rune{}, val[:i+1]...)
					secondPart := append([]rune{}, val[i+1:]...)
					firstPart = append(firstPart, '.')
					grid[j] = append(firstPart, secondPart...)
				}
			}
		}
	}

	return grid
}

func addCol(grid [][]rune, index int) [][]rune {
	for i := range grid {
		firstPart := append(grid[i], grid[i][:index+1]...)
		secondPart := append(grid[i], grid[i][index+1:]...)
		firstPart = append(firstPart, '.')
		grid[i] = append(firstPart, secondPart...)
		fmt.Println(grid)
	}

	return grid
}

func readFile(filepath string) [][]rune {
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

	grid := make([][]rune, 0)
	for i, v := range lines {
		grid = append(grid, []rune{})
		for _, val := range v {
			grid[i] = append(grid[i], val)
		}
	}

	return grid
}
