package main

import (
	"bufio"
	"fmt"
	"os"
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

	//create map containing each letter as key and its priority as value
	alphabet := make(map[string]int, 0)
	index := 1
	for i := 'a'; i <= 'z'; i++ {
		alphabet[string(i)] = index
		index++
	}
	for i := 'A'; i <= 'Z'; i++ {
		alphabet[string(i)] = index
		index++
	}

	var badges []string
	for i := 0; i < len(lines); i += 3 {
		for _, val := range lines[i] {
			if strings.Contains(lines[i+1], string(val)) && strings.Contains(lines[i+2], string(val)) {
				badges = append(badges, string(val))
				break
			}
		}
	}
	fmt.Println(len(badges))

	//find total
	total := 0
	for _, v := range badges {
		total += alphabet[v]
	}

	fmt.Println(total)
}
