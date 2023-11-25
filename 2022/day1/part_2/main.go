package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	elves := strings.Split(string(bytes), "\r\n")
	var elvesArr []int
	var totalCalories int
	for _, v := range elves {
		if v != "" {
			value, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			totalCalories += value
		} else {
			elvesArr = append(elvesArr, totalCalories)
			totalCalories = 0
		}
	}

	max := 0
	max2 := 0
	max3 := 0
	for _, v := range elvesArr {
		if v > max {
			max3 = max2
			max2 = max
			max = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}
	}
	fmt.Println(max + max2 + max3)
}
