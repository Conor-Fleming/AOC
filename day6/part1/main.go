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

	two := 1
	three := 2
	four := 3
	for one := 0; one < len(buff); one++ {
		fmt.Println(buff[one], buff[two], buff[three], buff[four])
		if buff[one] != buff[two] && buff[one] != buff[three] && buff[one] != buff[four] && buff[two] != buff[three] && buff[two] != buff[four] && buff[three] != buff[four] {
			fmt.Println("signal: ", one+4)
			break
		}
		two++
		three++
		four++
	}
}
