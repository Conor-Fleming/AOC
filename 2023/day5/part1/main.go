package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := readFile("../input.txt")

	//globals
	seeds := make([]int, 0)
	mapLines := make([][][]int, 0)
	locations := make([]int, 0)

	mapIndex := -1
	for i, v := range lines {
		if i == 0 {
			//getting seed values and converting to ints
			seedsStrings := strings.Fields(strings.Split(v, ":")[1])
			for _, v := range seedsStrings {
				seedInt, _ := strconv.Atoi(v)
				seeds = append(seeds, seedInt)
			}
			continue
		}

		if strings.Contains(v, ":") {
			mapIndex++
			continue
		}

		if v == "" {
			continue
		}

		fields := strings.Fields(v)
		lineDigits := make([]int, 0)
		for _, v := range fields {
			dig, _ := strconv.Atoi(v)
			lineDigits = append(lineDigits, dig)
		}

		for len(mapLines) <= mapIndex {
			mapLines = append(mapLines, make([][]int, 0))
		}

		mapLines[mapIndex] = append(mapLines[mapIndex], lineDigits)
	}

	//get location for each seed
	for _, v := range seeds {
		locations = append(locations, getSeedLocation(v, mapLines))
	}

	fmt.Println(slices.Min(locations))
}

func getSeedLocation(seed int, maps [][][]int) int {
	for _, v := range maps {
		for _, val := range v {
			if seed > val[1] && seed <= val[1]+val[2] {
				seed = seed + val[0] - val[1]
				break
			}
		}
	}

	return seed
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
