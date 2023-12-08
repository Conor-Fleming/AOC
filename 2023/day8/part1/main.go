package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readFile("../input.txt")
	nodes := make(map[string][]string)
	instructions := lines[0]
	for i, v := range lines {
		if i == 0 || v == "" {
			continue
		}

		parts := strings.Fields(v)
		nodes[parts[0]] = []string{parts[2][1:4], parts[3][:3]}
	}

	result := navigate("AAA", instructions, nodes, 0)

	fmt.Println(result)
}

func navigate(node string, instructions string, nodes map[string][]string, count int) int {
	if node == "ZZZ" {
		return count
	}

	instruction := instructions[count]
	instructions += string(instruction)
	count++
	if instruction == 'L' {
		return navigate(nodes[node][0], instructions, nodes, count)
	}

	if instruction == 'R' {
		return navigate(nodes[node][1], instructions, nodes, count)
	}

	return 0
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
