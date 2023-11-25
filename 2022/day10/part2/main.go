package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("../input.txt")
	x := 1
	cycle := -1

	for _, v := range lines {
		instruction := strings.Split(v, " ")
		//noop instruction
		if len(instruction) == 1 {
			cycle++
			cycle = checkCycle(cycle, x)
			continue
		} else {
			cycle++
			val, _ := strconv.Atoi(instruction[1])
			cycle = checkCycle(cycle, x)
			cycle++
			cycle = checkCycle(cycle, x)
			x += val
		}
	}
}

func checkCycle(cycle int, x int) int {
	if cycle > 39 {
		cycle = 0
		fmt.Print("\n")
	}
	if cycle == x-1 || cycle == x || cycle == x+1 {
		fmt.Printf("#")
	} else {
		fmt.Print(" ")
	}
	return cycle
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
