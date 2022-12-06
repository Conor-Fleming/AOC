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

	for i := range buff[:len(buff)-14] {
		test := []byte{}
		for j := 0; j < 14; j++ {
			test = append(test, buff[i+j])
		}
		result := removeDuplicateStr(test)
		if len(result) == len(test) && i >= 13 {
			fmt.Println(i + 14)
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
