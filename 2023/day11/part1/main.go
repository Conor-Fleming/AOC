package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	grid := readFile("../input.txt")
	grid = expandUnivRows(grid)
	grid = expandUnivCols(grid)

	for _, v := range grid {
		fmt.Println(string(v))
	}

}

func expandUnivCols(grid [][]rune) [][]rune {
	indexes := make([]int, 0)
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == '#' {
				break
			}

			if j == len(grid)-1 {
				indexes = append(indexes, i)
			}
		}

	}

	return addCol(grid, indexes)
}

func addCol(grid [][]rune, indexes []int) [][]rune {
	fmt.Println(indexes)
	for ext, v := range indexes {
		for i := range grid {
			firstPart := append([]rune{}, grid[i][:v+1+ext]...)
			secondPart := append([]rune{}, grid[i][v+1+ext:]...)
			firstPart = append(firstPart, '.')
			grid[i] = append(firstPart, secondPart...)
		}
	}

	return grid
}

func expandUnivRows(grid [][]rune) [][]rune {
	for i, v := range grid {
		if !strings.Contains(string(v), "#") {
			firstPart := append([][]rune{}, grid[:i+1]...)
			secondPart := append([][]rune{}, grid[i+1:]...)
			firstPart = append(firstPart, v)
			grid = append(firstPart, secondPart...)
		}
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
