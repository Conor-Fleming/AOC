package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	buff, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	for i := range buff[:len(buff)-4] {
		test := []byte{}
		for j := 0; j < 4; j++ {
			test = append(test, buff[i+j])
		}
		result := removeDuplicateStr(test)
		if len(result) == len(test) && i >= 3 {
			fmt.Println(i + 4)
			break
		}
	}
}

func removeDuplicateStr(byteSlice []byte) []byte {
	allKeys := make(map[byte]bool)
	list := []byte{}
	for _, item := range byteSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
