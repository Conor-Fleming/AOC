package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../guide.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	elves := strings.Split(string(bytes), "\n")
	var matches []string
	var match string
	for _, v := range elves {
		if v != "\n" {
			match += v
		}
		matches = append(matches, match)
	}

	/*total := 0
	for _, v := range matches {
		fmt.Println(v)
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
	}*/

	fmt.Println(matches)
	//fmt.Println(total)
}
