package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

func main() {
	lines := readFile("input.txt")
	grid := make([][]string, 0)
	for i, v := range lines {
		grid = append(grid, []string{})
		for _, val := range v {
			grid[i] = append(grid[i], string(val))
		}
	}

	fmt.Println("Part 1: ", part1(grid))
	fmt.Println("Part 2: ", part2(grid))
}

func part1(grid [][]string) int {
	var start point
	for i, v := range grid {
		for j, val := range v {
			if val == "^" {
				start = point{
					x: i,
					y: j,
				}
				break
			}
		}
	}

	visited := countPoints(start, grid)
	return visited
}

func part2(grid [][]string) int {
	var start point
	for i, v := range grid {
		for j, val := range v {
			if val == "^" {
				start = point{
					x: i,
					y: j,
				}
				break
			}
		}
	}

	visited := countPoints(start, grid)
	return visited
}

func countPoints(curr point, grid [][]string) int {
	record := make(map[point]bool)
	for {
	up:
		for { //up
			if _, ok := record[curr]; !ok {
				record[curr] = true
			}
			//checking top edge
			if curr.x == 0 {
				return len(record)
			}

			if grid[curr.x-1][curr.y] == "#" {
				goto right //go next direction (right)
			}

			curr = point{x: curr.x - 1, y: curr.y}
		}

	right:
		for { //right
			if _, ok := record[curr]; !ok {
				record[curr] = true
			}
			//right edge check
			if curr.y == len(grid[0])-1 {
				return len(record)
			}

			if grid[curr.x][curr.y+1] == "#" {
				goto down //go next direction (down)
			}

			curr = point{x: curr.x, y: curr.y + 1}
		}

	down:
		for { //down
			if _, ok := record[curr]; !ok {
				record[curr] = true
			}
			//bottom edge check
			if curr.x == len(grid)-1 {
				return len(record)
			}

			if grid[curr.x+1][curr.y] == "#" {
				goto left //go next direction (left)
			}

			curr = point{x: curr.x + 1, y: curr.y}
		}

	left:
		for { //left
			if _, ok := record[curr]; !ok {
				record[curr] = true
			}
			//left edge check
			if curr.y == 0 {
				return len(record)
			}

			if grid[curr.x][curr.y-1] == "#" {
				goto up //go next direction (up)
			}

			curr = point{x: curr.x, y: curr.y - 1}
		}
	}
}

func countLoops(curr point, grid [][]string) int {
	record := make(map[point]bool)
	for {
	up:
		for { //up
			//checking top edge
			if curr.x == 0 {
				return len(record)
			}

			if grid[curr.x-1][curr.y] == "#" {
				goto right //go next direction (right)
			}

			curr = point{x: curr.x - 1, y: curr.y}
		}

	right:
		for { //right
			//right edge check
			if curr.y == len(grid[0])-1 {
				return len(record)
			}

			if grid[curr.x][curr.y+1] == "#" {
				//curr = point{x: curr.x - 1, y: curr.y}
				goto down //go next direction (down)
			}

			curr = point{x: curr.x, y: curr.y + 1}
		}

	down:
		for { //down
			//bottom edge check
			if curr.x == len(grid)-1 {
				return len(record)
			}

			if grid[curr.x+1][curr.y] == "#" {
				//curr = point{x: curr.x, y: curr.y - 1}
				goto left //go next direction (left)
			}

			curr = point{x: curr.x + 1, y: curr.y}
		}

	left:
		for { //left
			//left edge check
			if curr.y == 0 {
				return len(record)
			}

			if grid[curr.x][curr.y-1] == "#" {
				//curr = point{x: curr.x - 1, y: curr.y}
				goto up //go next direction (up)
			}

			curr = point{x: curr.x, y: curr.y - 1}
		}
	}
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
