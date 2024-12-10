package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type file struct {
	id     int
	blocks int
	value  string
}

func main() {
	diskMap := getString("input.txt")

	fmt.Println("Part 1:", part1(diskMap))
}

func part1(diskMap string) int {
	fileSlice := make([]file, 0)
	position := 0
	for i, v := range diskMap {
		count, _ := strconv.Atoi(string(v))
		if i%2 == 0 {
			fileSlice = append(fileSlice, file{
				id:     position,
				blocks: count,
				value:  strconv.Itoa(position),
			})
			position++
		} else {
			free := file{}
			free.blocks = count
			free.value = "."

			fileSlice = append(fileSlice, free)
		}
	}

	for _, v := range fileSlice {
		for i := 0; i < v.blocks; i++ {
			fmt.Print(v.value)
		}
	}

	fmt.Println()
	return 0
}

func getString(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	diskMap := ""
	// Read the first line
	if scanner.Scan() {
		diskMap = scanner.Text()
	}

	return diskMap
}
