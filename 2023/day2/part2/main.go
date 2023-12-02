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
	hands := strings.Split(colonSpl[1], ";")
	re := regexp.MustCompile(`(\d+)\s+(green|red|blue)`)

	red := 0
	blue := 0
	green := 0

	for _, v := range hands {
		matches := re.FindAllString(v, -1)

		r, b, g := checkHandPossibile(matches)
		if r > red {
			red = r
		}
		if b > blue {
			blue = b
		}
		if g > green {
			green = g
		}
	}

	return red * blue * green
}

func checkHandPossibile(hand []string) (red, blue, green int) {
	red = 0
	blue = 0
	green = 0
	for _, v := range hand {
		color := strings.Split(v, " ")
		numb, _ := strconv.Atoi(color[0])
		switch color[1] {
		case "red":
			if numb > red {
				red = numb
				continue
			}
		case "blue":
			if numb > blue {
				blue = numb
				continue
			}
		case "green":
			if numb > green {
				green = numb
				continue
			}
		}
	}

	return red, blue, green
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
