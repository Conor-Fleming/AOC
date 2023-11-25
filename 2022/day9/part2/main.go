package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Coords struct {
	x int
	y int
}

func main() {
	direction, number := readFile("../input.txt")
	visited := make(map[Coords]bool, 0)
	rope := make([]Coords, 10)
	for i := range rope {
		rope[i] = Coords{0, 0}
	}

	visited[rope[len(rope)-1]] = true

	visited = moveRope(direction, number, visited, rope)

	fmt.Println(len(visited))
}

func moveRope(direction []string, number []int, visited map[Coords]bool, rope []Coords) map[Coords]bool {
	for i, v := range direction {
		iter := 0
		switch v {
		case "U":
			for iter < number[i] {
				rope[0].y++
				//move next knot
				for knot := range rope[:len(rope)-1] {
					rope[knot+1] = moveTail(rope[knot+1], rope[knot])
				}
				visited = ropeTailCoordsCheck(rope[9], visited)
				iter++
			}
		case "D":
			for iter < number[i] {
				rope[0].y--
				for knot := range rope[:len(rope)-1] {
					rope[knot+1] = moveTail(rope[knot+1], rope[knot])
				}
				visited = ropeTailCoordsCheck(rope[9], visited)
				iter++
			}
		case "L":
			for iter < number[i] {
				rope[0].x--
				for knot := range rope[:len(rope)-1] {
					rope[knot+1] = moveTail(rope[knot+1], rope[knot])
				}
				visited = ropeTailCoordsCheck(rope[9], visited)
				iter++
			}
		case "R":
			for iter < number[i] {
				rope[0].x++
				for knot := range rope[:len(rope)-1] {
					rope[knot+1] = moveTail(rope[knot+1], rope[knot])
				}
				visited = ropeTailCoordsCheck(rope[9], visited)
				iter++
			}
		}
	}
	return visited
}

func moveTail(tail Coords, head Coords) Coords {
	next := tail
	switch (Coords{head.x - tail.x, head.y - tail.y}) {
	case Coords{-2, 1}, Coords{-1, 2}, Coords{0, 2}, Coords{1, 2}, Coords{2, 1}, Coords{2, 2}, Coords{-2, 2}:
		next.y++
	}
	switch (Coords{head.x - tail.x, head.y - tail.y}) {
	case Coords{1, 2}, Coords{2, 1}, Coords{2, 0}, Coords{2, -1}, Coords{1, -2}, Coords{2, 2}, Coords{2, -2}:
		next.x++
	}
	switch (Coords{head.x - tail.x, head.y - tail.y}) {
	case Coords{-2, -2}, Coords{2, -1}, Coords{1, -2}, Coords{0, -2}, Coords{-1, -2}, Coords{-2, -1}, Coords{2, -2}:
		next.y--
	}
	switch (Coords{head.x - tail.x, head.y - tail.y}) {
	case Coords{-2, -2}, Coords{-1, -2}, Coords{-2, -1}, Coords{-2, -0}, Coords{-2, 1}, Coords{-1, 2}, Coords{-2, 2}:
		next.x--
	}
	return next
}

func ropeTailCoordsCheck(tail Coords, visited map[Coords]bool) map[Coords]bool {
	if _, ok := visited[tail]; !ok {
		visited[tail] = true
	}
	return visited
}

func readFile(filepath string) ([]string, []int) {
	//read contents of file to lines array
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	var number []int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, string(scanner.Text()[0]))
		val, _ := strconv.Atoi(scanner.Text()[2:])
		number = append(number, val)
	}

	return lines, number
}
