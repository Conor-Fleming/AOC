package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	games := readFile("../input.txt")

	total := 0
	for _, v := range games {
		total += checkGame(v)
	}

	fmt.Println(total)
}

func checkGame(game string) int {
	colonSpl := strings.Split(game, ":")
	id := strings.Split(colonSpl[0], " ")[1]
	hands := strings.Split(colonSpl[1], ";")
	re := regexp.MustCompile(`(\d+)\s+(green|red|blue)`)

	idInt, _ := strconv.Atoi(id)

	for _, v := range hands {
		matches := re.FindAllString(v, -1)

		if !checkHandPossibile(matches) {
			return 0
		}
	}

	return idInt
}

func checkHandPossibile(hand []string) bool {
	for _, v := range hand {
		color := strings.Split(v, " ")
		numb, _ := strconv.Atoi(color[0])
		switch color[1] {
		case "red":
			if numb > 12 {
				return false
			}
		case "blue":
			if numb > 14 {
				return false
			}
		case "green":
			if numb > 13 {
				return false
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

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
