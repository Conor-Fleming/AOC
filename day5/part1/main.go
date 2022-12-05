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
	arr1 := []rune{'B', 'G', 'S', 'C'}
	arr2 := []rune{'T', 'M', 'W', 'H', 'J', 'N', 'V', 'G'}
	arr3 := []rune{'M', 'Q', 'S'}
	arr4 := []rune{'B', 'S', 'L', 'T', 'W', 'N', 'M'}
	arr5 := []rune{'J', 'Z', 'F', 'T', 'V', 'G', 'W', 'P'}
	arr6 := []rune{'C', 'T', 'B', 'G', 'Q', 'H', 'S'}
	arr7 := []rune{'T', 'J', 'P', 'B', 'W'}
	arr8 := []rune{'G', 'D', 'C', 'Z', 'F', 'T', 'Q', 'M'}
	arr9 := []rune{'N', 'S', 'H', 'B', 'P', 'F'}

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
		v = strings.ReplaceAll(v, "move ", "")
		v = strings.ReplaceAll(v, "from ", "")
		v = strings.ReplaceAll(v, "to ", "")
		moveData := strings.Split(v, " ")

		quantityToMove, _ := strconv.Atoi(string(moveData[0]))
		movedFrom, _ := strconv.Atoi(string(moveData[1]))
		movedTo, _ := strconv.Atoi(string(moveData[2]))
		for i := 0; i < quantityToMove; i++ {
			to := fmt.Sprint("arr%v", movedTo)
			from := fmt.Sprintf("arr%v", movedFrom)
		}
	}
}
