package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	bet   int
	cards string
}

type hands []hand

func (h hands) Len() int {
	return len(h)
}

func (h hands) Less(i, j int) bool {
	return compareHands(h[i].cards, h[j].cards)
}

func (h hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Function to compare two hands based on their card values
func compareHands(cards1, cards2 string) bool {
	cardValues := "AKQT98765432J" // Define card values from highest to lowest
	for i := 0; i < len(cards1); i++ {
		index1 := strings.Index(cardValues, string(cards1[i]))
		index2 := strings.Index(cardValues, string(cards2[i]))

		if index1 != index2 {
			return index1 < index2 // Higher index means stronger card
		}
	}
	return false // If both hands are equal
}

func main() {
	lines := readFile("../input.txt")
	handsList := make(hands, 0)
	for _, v := range lines {
		handParts := strings.Fields(v)
		betValue, _ := strconv.Atoi(handParts[1])
		h := hand{
			bet:   betValue,
			cards: handParts[0],
		}

		handsList = append(handsList, h)
	}

	five, four, full, three, two, one, high := groupTypes(handsList)
	sorted := make([]hand, 0)

	sort.Sort(five)
	sorted = append(sorted, five...)
	sort.Sort(four)
	sorted = append(sorted, four...)
	sort.Sort(full)
	sorted = append(sorted, full...)
	sort.Sort(three)
	sorted = append(sorted, three...)
	sort.Sort(two)
	sorted = append(sorted, two...)
	sort.Sort(one)
	sorted = append(sorted, one...)
	sort.Sort(high)
	sorted = append(sorted, high...)

	winnings := 0
	for i := 0; i < len(sorted); i++ {
		rank := len(sorted) - i
		winnings += sorted[i].bet * rank
	}

	fmt.Println(winnings)
}

func groupTypes(handsList hands) (five, four, full, three, two, one, high hands) {
	five = make([]hand, 0)
	four = make([]hand, 0)
	full = make([]hand, 0)
	three = make([]hand, 0)
	two = make([]hand, 0)
	one = make([]hand, 0)
	high = make([]hand, 0)

	for _, v := range handsList {
		charCount := make(map[rune]int)
		for _, v := range v.cards {
			if _, ok := charCount[v]; !ok {
				charCount[v] = 1
				continue
			}

			charCount[v]++
		}

		tempMap := charCount
		max := 0
		maxKeys := make([]rune, 0)
		for k, v := range charCount {
			if v == 'J' {
				continue
			}
			if v > max {
				max = v
				maxKeys = []rune{k}
			} else if v == max {
				maxKeys = append(maxKeys, k)
			}
		}

		maxK := maxKeys[0]
		if len(maxKeys) > 1 {
			maxK = maxKeys[0]
			for _, v := range maxKeys {
				if v > maxK {
					maxK = v
				}
			}
		}

		for k := range tempMap {
			if k == 'J' {
				tempMap[maxK] += tempMap[k]
				delete(tempMap, k)
			}
		}

		switch len(tempMap) {
		case 1:
			five = append(five, v)
		case 2:
			found := false
			for _, val := range charCount {
				if val == 4 {
					found = true
				}
			}

			if found {
				four = append(four, v)
			} else {
				full = append(full, v)
			}

		case 3:
			found := false
			for _, val := range charCount {
				if val == 3 {
					found = true
				}
			}

			if found {
				three = append(three, v)
			} else {
				two = append(two, v)
			}
		case 4:
			one = append(one, v)
		default:
			high = append(high, v)
		}
	}

	return five, four, full, three, two, one, high
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
