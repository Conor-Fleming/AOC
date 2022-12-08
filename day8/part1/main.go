package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := readFile("../input.txt")
	total := 0

	for i, v := range lines {
		for j, val := range v {
			//outter edge of matrix will always be visable
			if i == 0 || j == 0 || i == len(lines)-1 || j == len(lines[0])-1 {
				total++
				continue
			}
			//if either func returns true add one to total
			if checkRow(lines, byte(val), i, j) || checkColumn(lines, byte(val), i, j) {
				total++
			}
		}
	}

	fmt.Println(total)
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
