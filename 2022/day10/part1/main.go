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
	cycle := 0
	var signals []int
	for _, v := range lines {
		instruction := strings.Split(v, " ")
		//noop instruction
		if len(instruction) == 1 {
			cycle++
			signals = checkCycle(cycle, x, signals)
			continue
		} else {
			cycle++
			val, _ := strconv.Atoi(instruction[1])
			signals = checkCycle(cycle, x, signals)
			cycle++
			signals = checkCycle(cycle, x, signals)
			x += val
		}
	}

	result := 0
	for _, v := range signals {
		result += v
	}
	fmt.Println(result)
}

func checkCycle(cycle int, x int, signals []int) []int {
	if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
		signals = append(signals, x*cycle)
	}
	return signals
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
