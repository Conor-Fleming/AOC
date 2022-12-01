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
	number := 0
	for i, v := range elvesArr {
		if v > max {
			max = v
			number = i
		}
	}
	fmt.Println(number, ":", max)

}
