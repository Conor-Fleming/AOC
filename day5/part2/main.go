package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*

	[G]         [P]         [M]
	[V]     [M] [W] [S]     [Q]
	[N]     [N] [G] [H]     [T] [F]
	[J]     [W] [V] [Q] [W] [F] [P]
[C] [H]     [T] [T] [G] [B] [Z] [B]
[S] [W] [S] [L] [F] [B] [P] [C] [H]
[G] [M] [Q] [S] [Z] [T] [J] [D] [S]
[B] [T] [M] [B] [J] [C] [T] [G] [N]
1   2   3   4   5   6   7   8   9

*/

func main() {
	crates := [][]rune{
		{'B', 'G', 'S', 'C'},
		{'T', 'M', 'W', 'H', 'J', 'N', 'V', 'G'},
		{'M', 'Q', 'S'},
		{'B', 'S', 'L', 'T', 'W', 'N', 'M'},
		{'J', 'Z', 'F', 'T', 'V', 'G', 'W', 'P'},
		{'C', 'T', 'B', 'G', 'Q', 'H', 'S'},
		{'T', 'J', 'P', 'B', 'W'},
		{'G', 'D', 'C', 'Z', 'F', 'T', 'Q', 'M'},
		{'N', 'S', 'H', 'B', 'P', 'F'},
	}

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

	for _, v := range lines {
		instructions := strings.ReplaceAll(v, "move ", "")
		instructions = strings.ReplaceAll(instructions, "from ", "")
		instructions = strings.ReplaceAll(instructions, "to ", "")
		moveData := strings.Split(instructions, " ")

		quantityToMove, _ := strconv.Atoi(string(moveData[0]))
		movedFrom, _ := strconv.Atoi(string(moveData[1]))
		movedFrom--
		movedTo, _ := strconv.Atoi(string(moveData[2]))
		movedTo--

		n := len(crates[movedFrom])

		crates[movedTo] = append(crates[movedTo], crates[movedFrom][n-quantityToMove:]...)
		crates[movedFrom] = crates[movedFrom][:n-quantityToMove]
	}
	for _, v := range crates {
		fmt.Println(string(v))
	}
}
