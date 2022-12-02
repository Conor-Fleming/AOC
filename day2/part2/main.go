package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("../guide.txt")
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
		fmt.Println(v)
		switch v {
		case "A X":
			//lose 0
			//rock beats scissors(3)
			total += 3
		case "A Y":
			//draw 3
			//rock draws rock(1)
			total += 4
		case "A Z":
			//win 6
			//rock loses (paper: 2)
			total += 8
		case "B X":
			//lose 0
			total += 1
		case "B Y":
			//draw 3
			total += 5
		case "B Z":
			//win 6
			total += 9
		case "C X":
			//lose 0
			total += 2
		case "C Y":
			//draw 3
			total += 6
		case "C Z":
			//win 6
			total += 7
		default:
			continue
		}
	}

	fmt.Println(total)
}
