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
	head := Coords{0, 0}
	tail := Coords{0, 0}
	visited[tail] = true

	for i, v := range direction {
		num := number[i]
		switch v {
		case "U":
			for x := 0; x < num; x++ {
				head.y++
				if head.y-tail.y > 1 {
					if head.x > tail.x {
						tail.x++
					}
					if head.x < tail.x {
						tail.x--
					}
					tail.y++
					visited = tailCoordsCheck(tail, visited)
				}
			}
		case "D":
			for x := 0; x < num; x++ {
				head.y--
				if tail.y-head.y > 1 {
					if head.x > tail.x {
						tail.x++

					}
					if head.x < tail.x {
						tail.x--
					}
					tail.y--
					visited = tailCoordsCheck(tail, visited)
				}
			}

		case "L":
			for x := 0; x < num; x++ {
				head.x--
				if tail.x-head.x > 1 {
					if head.y > tail.y {
						tail.y++
					}
					if head.y < tail.y {
						tail.y--
					}
					tail.x--
					visited = tailCoordsCheck(tail, visited)
				}
			}

		case "R":
			for x := 0; x < num; x++ {
				head.x++
				if head.x-tail.x > 1 {
					if head.y > tail.y {
						tail.y++
					}
					if head.y < tail.y {
						tail.y--
					}
					tail.x++
					visited = tailCoordsCheck(tail, visited)
				}
			}

		}

	}
	fmt.Println(len(visited))

}

func tailCoordsCheck(tail Coords, visited map[Coords]bool) map[Coords]bool {
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
