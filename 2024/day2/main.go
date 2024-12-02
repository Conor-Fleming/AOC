package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reports := make([][]int, 0)
	for _, v := range readFile("input.txt") {
		temp := strings.Fields(v)
		report := make([]int, 0)
		for _, v := range temp {
			number, _ := strconv.Atoi(v)
			report = append(report, number)
		}
		reports = append(reports, report)
	}
	fmt.Println("Part 1: ", part1(reports))
	fmt.Println("Part 2: ", part2(reports))
}

func part1(reports [][]int) int {
	safeCount := 0
	for _, v := range reports {
		if isSafe(v) {
			safeCount++
		}
	}

	return safeCount
}

func part2(reports [][]int) int {
	safeCount := 0
	for _, v := range reports {
		if isSafe(v) {
			safeCount++
			continue
		}

		for i := 0; i < len(v); i++ {
			new := append([]int(nil), v[:i]...) // Create a fresh slice
			new = append(new, v[i+1:]...)       // Append the rest of the slice
			if isSafe(new) {
				safeCount++
				break
			}
		}
	}

	return safeCount
}

func isSafe(report []int) bool {
	dupes := make(map[int]int, 0)
	for _, v := range report {
		if _, ok := dupes[v]; ok {
			return false
		}
		dupes[v] = 1
	}

	if !isSortedAsc(report) && !isSortedDesc(report) {
		return false
	}

	for i := 1; i < len(report); i++ {
		if math.Abs(float64(report[i-1]-report[i])) > 3 {
			return false
		}
	}

	return true
}

func isSortedDesc(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] < slice[i] {
			return false
		}
	}
	return true
}

func isSortedAsc(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
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

	//var words []string
	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Split the line into words
		//lineWords := strings.Fields(scanner.Text())
		lines = append(lines, scanner.Text())

		// Append each word to the slice
		//words = append(words, lineWords...)
	}

	return lines
}
