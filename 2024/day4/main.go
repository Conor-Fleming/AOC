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

	fmt.Println("Part 1: ", part1(lines))
	fmt.Println("Part 1: ", part2(lines))
}

func part1(lines []string) int {
	points := make([]point, 0)
	for i, v := range lines {
		for j, val := range v {
			if val == 'X' {
				points = append(points, point{
					x: i,
					y: j,
				})
			}
		}
	}

	count := 0
	for _, v := range points {

		//Up
		if v.x >= 3 {
			if "XMAS" == string(lines[v.x][v.y])+string(lines[v.x-1][v.y])+string(lines[v.x-2][v.y])+string(lines[v.x-3][v.y]) {
				count++
			}
		}
		//Down
		if v.x < len(lines)-3 {
			if "XMAS" == string(lines[v.x][v.y])+string(lines[v.x+1][v.y])+string(lines[v.x+2][v.y])+string(lines[v.x+3][v.y]) {
				count++
			}
		}
		//Left
		if v.y >= 3 {
			if "XMAS" == string(lines[v.x][v.y])+string(lines[v.x][v.y-1])+string(lines[v.x][v.y-2])+string(lines[v.x][v.y-3]) {
				count++
			}
		}
		//Right
		if v.y < len(lines[v.x])-3 {
			if "XMAS" == string(lines[v.x][v.y])+string(lines[v.x][v.y+1])+string(lines[v.x][v.y+2])+string(lines[v.x][v.y+3]) {
				count++
			}
		}
		//DiagnalUpRight
		if v.x >= 3 && v.y < len(lines[v.x])-3 {
			if "XMAS" == string(lines[v.x][v.y])+string(lines[v.x-1][v.y+1])+string(lines[v.x-2][v.y+2])+string(lines[v.x-3][v.y+3]) {
				count++
			}
		}
		//DiagnalUpLeft
		if v.x >= 3 && v.y >= 3 {
			if "XMAS" == string(lines[v.x][v.y])+string(lines[v.x-1][v.y-1])+string(lines[v.x-2][v.y-2])+string(lines[v.x-3][v.y-3]) {
				count++
			}
		}
		//DiagnalDownRight
		if v.x < len(lines)-3 && v.y < len(lines[v.x])-3 {
			if "XMAS" == string(lines[v.x][v.y])+string(lines[v.x+1][v.y+1])+string(lines[v.x+2][v.y+2])+string(lines[v.x+3][v.y+3]) {
				count++
			}
		}
		//DiagnalDownLeft
		if v.x < len(lines)-3 && v.y >= 3 {
			if "XMAS" == string(lines[v.x][v.y])+string(lines[v.x+1][v.y-1])+string(lines[v.x+2][v.y-2])+string(lines[v.x+3][v.y-3]) {
				count++
			}
		}
	}

	return count
}

func part2(lines []string) int {
	points := make([]point, 0)
	for i, v := range lines {
		for j, val := range v {
			if val == 'A' {
				points = append(points, point{
					x: i,
					y: j,
				})
			}
		}
	}

	count := 0
	for _, v := range points {
		//DiagnalUpRight DiagnalUpLeft DiagnalDownRight DiagnalDownLeft
		if (v.x >= 1 && v.y < len(lines[v.x])-1) && (v.x >= 1 && v.y >= 1) && (v.x < len(lines)-1 && v.y < len(lines[v.x])-1) && (v.x < len(lines)-1 && v.y >= 1) {
			//M M
			// A
			//S S
			if (string(lines[v.x-1][v.y-1]) == "M" && string(lines[v.x+1][v.y+1]) == "S") && (string(lines[v.x-1][v.y+1]) == "M" && string(lines[v.x+1][v.y-1]) == "S") {
				count++
			}

			//S S
			// A
			//M M
			if (string(lines[v.x-1][v.y-1]) == "S" && string(lines[v.x+1][v.y+1]) == "M") && (string(lines[v.x-1][v.y+1]) == "S" && string(lines[v.x+1][v.y-1]) == "M") {
				count++
			}

			//M S
			// A
			//M S
			if (string(lines[v.x-1][v.y-1]) == "M" && string(lines[v.x+1][v.y+1]) == "S") && (string(lines[v.x-1][v.y+1]) == "S" && string(lines[v.x+1][v.y-1]) == "M") {
				count++
			}

			//S M
			// A
			//S M
			if (string(lines[v.x-1][v.y-1]) == "S" && string(lines[v.x+1][v.y+1]) == "M") && (string(lines[v.x-1][v.y+1]) == "M" && string(lines[v.x+1][v.y-1]) == "S") {
				count++
			}
		}
	}

	return count
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
