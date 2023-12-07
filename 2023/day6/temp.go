package main

import(
	"strings"
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

type hand struct {
	bet int
	cards string
}

var valueMap = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

func main(){
	lines := readFile("input.txt")
	handsList := make([]hand, 0)
	for _, v := range lines{
		handParts := strings.Fields(v)
		betValue, _ := strconv.Atoi(handParts[1])
		h := hand{
			bet: betValue,
			cards: handParts[0],
		}

		handsList = append(handsList, h)
	}

	five, four, full, three, two, one, high := groupTypes(handsList)
	sorted := make([]hand, 0)

	sorted = append(sorted, sortTypes(five)...)
	sorted = append(sorted, sortTypes(four)...)
	sorted = append(sorted, sortTypes(full)...)
	sorted = append(sorted, sortTypes(three)...)
	sorted = append(sorted, sortTypes(two)...)
	sorted = append(sorted, sortTypes(one)...)
	sorted = append(sorted, sortTypes(high)...)

	winnings := 0
	for i := 0; i < len(sorted); i++{
		rank := len(sorted) - i
		winnings += sorted[i].bet * rank
	}

	fmt.Println(winnings)
}

func sortTypes(groupType []hand) []hand {
	sort.Slice(groupType, func(i, j int) bool {
		for k := 0; k < len(groupType[i].cards); k++{
			if groupType[i].cards[k] != groupType[j].cards[k] {
				if groupType[i].cards[k] >= 'A' && groupType[i].cards[k] <= 'Z'{
					if groupType[j].cards[k] >= 'A' && groupType[j].cards[k] <= 'Z'{
						return groupType[i].cards[k] > groupType[j].cards[k]
					} else{
						return true
					}
				} else {
					if groupType[j].cards[k] >= 'A' && groupType[j].cards[k] <= 'Z'{
						return false
					} else {
						return groupType[i].cards[k] > groupType[j].cards[k]
					}
				}
				//return valueMap[groupType[i].cards[k]] < valueMap[groupType[j].cards[k]]
			}
		}
		return false
	})

	return groupType
}

func groupTypes(handsList []hand) (five, four, full, three, two, one, high []hand) {
	five = make([]hand, 0)
	four = make([]hand, 0)
	full = make([]hand, 0)
	three = make([]hand, 0)
	two = make([]hand, 0)
	one = make([]hand, 0)
	high = make([]hand, 0)

	for _, v := range handsList{
		charCount := make(map[rune]int)
		for _, v := range v.cards{
			if _, ok := charCount[v]; !ok{
				charCount[v] = 1
				continue
			}

			charCount[v]++
		}

		switch len(charCount) {
		case 1:
			five = append(five, v)
			break
		case 2:
			found := false
			for _, val := range charCount{
				if val == 4 {
					found = true
				}
			}

			if found{
				four = append(four, v)
			} else{
				full = append(full, v)
			}
			break
		case 3:
			found := false
			for _, val := range charCount{
				if val == 3 {
					found = true
				}
			}

			if found{
				three = append(three, v)
			} else{
				two = append(two, v)
			}
			break
		case 4:	
			one = append(one, v)
			break
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
