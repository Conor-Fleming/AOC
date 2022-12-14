package main

import "fmt"

func main() {
	monkeys := [][]int{
		{56, 52, 58, 96, 70, 75, 72},
		{75, 58, 86, 80, 55, 81},
		{73, 68, 73, 90},
		{72, 89, 55, 51, 59},
		{76, 76, 91},
		{88},
		{64, 63, 56, 50, 77, 55, 55, 86},
		{79, 58},
	}

	var monk0 int
	var monk1 int
	var monk2 int
	var monk3 int
	var monk4 int
	var monk5 int
	var monk6 int
	var monk7 int

	for x := 0; x < 20; x++ {
		for i, v := range monkeys {
			if len(v) == 0 {
				fmt.Println("empty", i)
				continue
			}
			for _, val := range v {
				if i == 0 {
					monk0++
					item := (val * 17) / 3
					monkeys[i] = monkeys[i][1:]
					if item%11 == 0 {
						monkeys[2] = append(monkeys[2], item)
					} else {
						monkeys[3] = append(monkeys[3], item)
					}
				}
				if i == 1 {
					monk1++
					item := (val + 7) / 3
					monkeys[i] = monkeys[i][1:]
					if item%3 == 0 {
						monkeys[6] = append(monkeys[6], item)
					} else {
						monkeys[5] = append(monkeys[5], item)
					}
				}
				if i == 2 {
					monk2++
					item := (val * val) / 3
					monkeys[i] = monkeys[i][1:]
					if item%5 == 0 {
						monkeys[1] = append(monkeys[1], item)
					} else {
						monkeys[7] = append(monkeys[7], item)
					}
				}
				if i == 3 {
					monk3++
					item := (val + 1) / 3
					monkeys[i] = monkeys[i][1:]
					if item%7 == 0 {
						monkeys[2] = append(monkeys[2], item)
					} else {
						monkeys[7] = append(monkeys[7], item)
					}
				}
				if i == 4 {
					monk4++
					item := (val * 3) / 3
					monkeys[i] = monkeys[i][1:]

					if item%19 == 0 {
						monkeys[0] = append(monkeys[0], item)
					} else {
						monkeys[3] = append(monkeys[3], item)
					}
				}
				if i == 5 {
					monk5++
					item := (val + 4) / 3
					monkeys[i] = monkeys[i][1:]
					if item%2 == 0 {
						monkeys[6] = append(monkeys[6], item)
					} else {
						monkeys[4] = append(monkeys[4], item)
					}
				}
				if i == 6 {
					monk6++
					item := (val + 8) / 3
					monkeys[i] = monkeys[i][1:]
					if item%13 == 0 {
						monkeys[4] = append(monkeys[4], item)
					} else {
						monkeys[0] = append(monkeys[0], item)
					}
				}
				if i == 7 {
					monk7++
					item := (val + 6) / 3
					monkeys[i] = monkeys[i][1:]
					if item%17 == 0 {
						monkeys[1] = append(monkeys[1], item)
					} else {
						monkeys[5] = append(monkeys[5], item)
					}
				}
			}
		}
	}
	fmt.Println(monk0, monk1, monk2, monk3, monk4, monk5, monk6, monk7)
}
