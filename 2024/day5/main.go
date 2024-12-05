package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rules, updates := readFile("rules.txt", "updates.txt")

	fmt.Println("Part 1: ", part1(rules, updates))
	fmt.Println("Part 2: ", part2(rules, updates))
}

func part1(rules, updates []string) int {
	count := 0
	valid := true
	for _, update := range updates {
		for _, rule := range rules {
			if strings.Contains(update, rule[:2]) && strings.Contains(update, rule[3:]) {
				if strings.Index(update, rule[:2]) > strings.Index(update, rule[3:]) {
					valid = false
					break
				}
			}
		}
		if valid {
			middle := update[len(update)/2-1 : len(update)/2+1]

			value, _ := strconv.Atoi(middle)
			count += value
		}
		valid = true
	}

	return count
}

func part2(rules, updates []string) int {
	count := 0
	fixed := false
	for _, update := range updates {
		for i := 0; i < len(rules); i++ {
			if strings.Contains(update, rules[i][:2]) && strings.Contains(update, rules[i][3:]) {
				if strings.Index(update, rules[i][:2]) > strings.Index(update, rules[i][3:]) {
					update = strings.Replace(update, rules[i][:2], rules[i][3:], 1)
					update = strings.Replace(update, rules[i][3:], rules[i][:2], 1)
					fixed = true
					i = 0
				}
			}
		}
		if fixed {
			middle := update[len(update)/2-1 : len(update)/2+1]

			value, _ := strconv.Atoi(middle)
			count += value
		}
		fixed = false
	}

	return count
}

func readFile(filepath1, filepath2 string) ([]string, []string) {
	//read contents of file to lines array
	file1, err := os.Open(filepath1)
	file2, err := os.Open(filepath2)
	if err != nil {
		panic(err)
	}
	defer file1.Close()
	defer file2.Close()

	var rules []string
	var updates []string
	scanner := bufio.NewScanner(file1)
	for scanner.Scan() {
		rules = append(rules, scanner.Text())
	}

	scanner = bufio.NewScanner(file2)
	for scanner.Scan() {
		updates = append(updates, scanner.Text())
	}

	return rules, updates
}
