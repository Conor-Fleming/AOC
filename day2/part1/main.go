package main

import (
	"bufio"
	"fmt"
	"os"
<<<<<<< HEAD
"bufio"
=======
>>>>>>> 0aceeca9c5f8110f9677b1a14717c1df6d4ebffa
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
<<<<<<< HEAD
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	
	total := 0
	for _, v := range lines {
=======
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	total := 0
	for _, v := range lines {
		fmt.Println(v)
>>>>>>> 0aceeca9c5f8110f9677b1a14717c1df6d4ebffa
		switch v {
		case "A X":
			total += 4
		case "A Y":
			total += 8
		case "A Z":
			total += 3
		case "B X":
			total++
		case "B Y":
			total += 5
		case "B Z":
			total += 9
		case "C X":
			total += 7
		case "C Y":
			total += 2
		case "C Z":
			total += 6
		default:
			continue
		}
	}

	fmt.Println(total)
}
