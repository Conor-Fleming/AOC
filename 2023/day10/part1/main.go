package main

import (
	"bufio"
	"fmt"
	"os"
)

type tile struct {
	coord coord
	value rune
}

type coord struct {
	i, j int
}

func main() {
	grid := readFile("../input.txt")
	start := findStart(grid)
	loop := getLoop(start, grid)

	fmt.Println(len(loop) / 2)
}

func getLoop(start tile, grid [][]rune) []tile {
	for _, v := range "|-JLF7" {
		start.value = v
		result := checkLoop(start, grid)
		if result == nil {
			continue
		}
	}

}

func checkLoop(s tile, grid [][]rune) []tile {

}

func findStart(grid [][]rune) tile {
	for i, v := range grid {
		for j, val := range v {
			if val == 'S' {
				startTile := tile{
					coord: coord{
						i: i,
						j: j,
					},
				}

				return startTile
			}
		}
	}

	//should never reach this line but if no start position will return empty tile obj
	return tile{}
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
