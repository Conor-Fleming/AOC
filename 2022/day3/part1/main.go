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

	//get two halves of each string and find common value
	var items []string
	for _, v := range lines {
		firstHalf := v[0 : len(v)/2]
		secondHalf := v[len(v)/2:]
		for _, val := range firstHalf {
			if strings.Contains(secondHalf, string(val)) {
				items = append(items, string(val))
				break
			}
		}
	}

	//find total
	total := 0
	for _, v := range items {
		total += alphabet[v]
	}

	fmt.Println(total)
}
