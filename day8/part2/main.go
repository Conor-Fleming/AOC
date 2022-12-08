package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readFile("../input.txt")
	max := 0

	for i, v := range lines {
		for j, val := range v {
			result := checkView(lines, byte(val), i, j)
			if result > max {
				max = result
			}
		}
	}

	fmt.Println(max)
}

func checkView(lines []string, val byte, row, col int) int {
	countright := 0
	countleft := 0
	countup := 0
	countdown := 0
	for columnright := col + 1; columnright < len(lines[row]); columnright++ {
		countright++
		if val <= lines[row][columnright] {
			break
		}
	}
	for columnleft := col - 1; columnleft >= 0; columnleft-- {
		countleft++
		if val <= lines[row][columnleft] {
			break
		}
	}
	for rowdown := row + 1; rowdown < len(lines); rowdown++ {
		countdown++
		if val <= lines[rowdown][col] {
			break
		}
	}
	for rowup := row - 1; rowup >= 0; rowup-- {
		countup++
		if val <= lines[rowup][col] {
			break
		}
	}
	return countup * countdown * countleft * countright
}

// check the row
func checkRow(lines []string, val byte, row, col int) bool {
	for columnright := col + 1; columnright < len(lines[row]); columnright++ {
		if val <= lines[row][columnright] {
			for columnleft := col - 1; columnleft >= 0; columnleft-- {
				if val <= lines[row][columnleft] {
					return false
				}
			}
		}
	}
	return true
}

// check the column
func checkColumn(lines []string, val byte, row, col int) bool {
	for rowdown := row + 1; rowdown < len(lines); rowdown++ {
		if val <= lines[rowdown][col] {
			for rowup := row - 1; rowup >= 0; rowup-- {
				if val <= lines[rowup][col] {
					return false
				}
			}
		}
	}
	return true
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
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
