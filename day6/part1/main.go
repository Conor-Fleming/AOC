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

	for _, v := range buff {
		test := []byte{}
		for i := 0; i < 14; i++ {
			test = append(test, buff[i])
			fmt.Print(len(test))
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
