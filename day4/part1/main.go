package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//read contents of file to lines array
	file, err := os.Open("../input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	total := 0
	for _, v := range lines {
		pair := strings.Split(v, ",")
		range1 := strings.Split(pair[0], "-")
		range2 := strings.Split(pair[1], "-")

		minPair1, _ := strconv.Atoi(range1[0])
		minPair2, _ := strconv.Atoi(range2[0])
		maxPair1, _ := strconv.Atoi(range1[1])
		maxPair2, _ := strconv.Atoi(range2[1])

		if minPair1 <= minPair2 && maxPair1 >= maxPair2 || minPair2 <= minPair1 && maxPair2 >= maxPair1 {
			total++
		}
	}

	fmt.Println(total)
}
