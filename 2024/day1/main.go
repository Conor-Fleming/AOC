package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	left := make([]int, 0)
	right := make([]int, 0)
	for i, v := range readFile("input.txt") {
		v, _ := strconv.Atoi(v)
		if i%2 == 0 {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	fmt.Println("Part 1: ", part1(left, right))
	fmt.Println("Part 2: ", part2(left, right))
}

func part1(left, right []int) int {
	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})
	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	distances := make([]float64, 0)
	for i, v := range left {
		distances = append(distances, math.Abs(float64(v-right[i])))
	}

	sum := 0
	for _, v := range distances {
		sum += int(v)
	}

	return sum
}

func part2(left, right []int) int {
	similarities := make([]int, 0)
	for _, v := range left {
		repeat := 0
		for _, val := range right {
			if v == val {
				repeat++
			}
		}
		similarities = append(similarities, (repeat * v))
	}

	sum := 0
	for _, v := range similarities {
		sum += v
	}

	return sum
}

func readFile(filepath string) []string {
	//read contents of file to lines array
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var words []string
	//var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Split the line into words
		lineWords := strings.Fields(scanner.Text())
		//ines = append(lines, scanner.Text())

		// Append each word to the slice
		words = append(words, lineWords...)
	}

	return words
}
