package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type race struct {
	time int
	dist int
}

func main() {
	races := make([]race, 0)

	lines := readFile("../input.txt")
	timeVals := strings.ReplaceAll(strings.Split(lines[0], ":")[1], " ", "")
	distanceVals := strings.ReplaceAll(strings.Split(lines[1], ":")[1], " ", "")

	time, _ := strconv.Atoi(timeVals)
	dist, _ := strconv.Atoi(distanceVals)

	r := race{
		time: time,
		dist: dist,
	}

	races = append(races, r)

	recordBeaters := make([]int, 0)
	for _, v := range races {
		count := 0
		for i := 1; i <= v.time; i++ {
			travTime := v.time - i
			newDist := travTime * i

			if newDist > v.dist {
				count++
			}
		}

		recordBeaters = append(recordBeaters, count)
	}

	result := 1
	for _, v := range recordBeaters {
		result *= v
	}

	fmt.Println(result)
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
