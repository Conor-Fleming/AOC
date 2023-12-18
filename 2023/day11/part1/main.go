package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type galaxyPair struct {
	first, second galaxy
}

type galaxy struct {
	val  string
	x, y int
}

func main() {
	grid := readFile("../input.txt")
	grid = expandUnivRows(grid)
	grid = expandUnivCols(grid)

	//get slice of pairs of galaxies
	pairs := numberAndPair(grid)
	fmt.Println(len(pairs))

}

func numberAndPair(grid [][]string) map[galaxyPair]bool {
	count := 1
	pairs := make(map[galaxyPair]bool)
	gals := make([]galaxy, 0)
	for i, v := range grid {
		for j, val := range v {
			if val == "#" {
				grid[i][j] = strconv.Itoa(count)
				gals = append(gals, galaxy{grid[i][j], i, j})
				count++
			}
		}
	}

	for i, v := range gals {
		for j := i + 1; j < len(gals); j++ {
			pair := galaxyPair{v, gals[j]}
			if _, ok := pairs[pair]; !ok {
				pairs[pair] = true
			}
		}
	}

	return pairs
}

func expandUnivCols(grid [][]string) [][]string {
	indexes := make([]int, 0)
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == "#" {
				break
			}

			if j == len(grid)-1 {
				indexes = append(indexes, i)
			}
		}

	}

	return addCol(grid, indexes)
}

func addCol(grid [][]string, indexes []int) [][]string {
	for ext, v := range indexes {
		for i := range grid {
			firstPart := append([]string{}, grid[i][:v+1+ext]...)
			secondPart := append([]string{}, grid[i][v+1+ext:]...)
			firstPart = append(firstPart, ".")
			grid[i] = append(firstPart, secondPart...)
		}
	}

	return grid
}

func expandUnivRows(grid [][]string) [][]string {
	for i, v := range grid {
		if !slices.Contains(v, "#") {
			firstPart := append([][]string{}, grid[:i+1]...)
			secondPart := append([][]string{}, grid[i+1:]...)
			firstPart = append(firstPart, v)
			grid = append(firstPart, secondPart...)
		}
	}

	return grid
}

func readFile(filepath string) [][]string {
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

	grid := make([][]string, 0)
	for i, v := range lines {
		grid = append(grid, []string{})
		for _, val := range v {
			grid[i] = append(grid[i], string(val))
		}
	}

	return grid
}
