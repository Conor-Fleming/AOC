package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	lines := readFile("input.txt")
	nodes := make(map[string][]string)
	instructions := lines[0]
	for i, v := range lines {
		if i == 0 || v == "" {
			continue
		}

		parts := strings.Fields(v)
		nodes[parts[0]] = []string{parts[2][1:4], parts[3][:3]}
	}

	startNodes := make([]string, 0)
	for k := range nodes{
		if k[2] == 'A' {
			startNodes = append(startNodes, k)
		}
	}

	counts := make([]int, 0)
	for _, v := range startNodes{
		count := 0
		counts = append(counts, navigate(v, instructions, nodes, count))
	}

	result := LCM(counts[0], counts[1], counts[2:]...)

	fmt.Println(result)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func navigate(node string, instructions string, nodes map[string][]string, count int) int {
	if node[2] == 'Z' {
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
