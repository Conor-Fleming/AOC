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
	sequences := make([][]int, 0)
	sequence := make([]int, 0)

	for _, v := range lines {
		sequence = nil
		for _, val := range strings.Fields(v) {
			converted, _ := strconv.Atoi(val)
			sequence = append(sequence, converted)
		}
		sequences = append(sequences, sequence)
	}

	result := 0
	for _, v := range sequences {
		storage := make([][]int, 0)
		result += findNext(unpack(v, storage))
	}

	fmt.Println(result)
}

func findNext(seqs [][]int) int {

	for i := len(seqs) - 1; i > 0; i-- {
		seqs[i-1] = append(seqs[i-1], seqs[i][len(seqs[i])-1]+seqs[i-1][len(seqs[i])-1])
	}

	return seqs[0][len(seqs[0])-1]
}

func allZero(slice []int) bool {
	for _, v := range slice {
		if v != 0 {
			return false
		}
	}

	return true
}

func unpack(diff []int, storage [][]int) [][]int {
	if allZero(diff) {
		diff = append(diff, 0)
		storage = append(storage, diff)
		return storage
	}
	storage = append(storage, diff)
	temp := make([]int, 0)
	for i := 0; i < len(diff)-1; i++ {
		temp = append(temp, diff[i+1]-diff[i])
	}

	return unpack(temp, storage)
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
