package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Tree struct {
	Size     int
	Name     string
	Dir      bool
	Parent   *Tree
	Children []*Tree
}

func (t *Tree) addDirectory(name string) {
	t.Children = append(t.Children, &Tree{
		Name:     name,
		Dir:      true,
		Parent:   t,
		Children: []*Tree{},
	})
}

func (t *Tree) addFile(name string, size int) {
	t.Children = append(t.Children, &Tree{
		Name:   name,
		Dir:    false,
		Parent: t,
	})

	t.countSize(size)
}

func (t *Tree) countSize(size int) {
	t.Size += size
	if t.Parent != nil {
		t.Parent.countSize(size)
	}
}

func (t *Tree) getDirectory(name string) *Tree {
	for _, v := range t.Children {
		if v.Name == name && v.Dir {
			return v
		}
	}
	return nil
}

func main() {
	lines := readFile("../input.txt")

	//setting root of tree
	root := &Tree{Name: "/", Children: []*Tree{}}
	cur := root

	//create filesystem tree
	for i := 0; i < len(lines); i++ {
		value := lines[i]
		switch {
		// return to root directory
		// sets current directory to root
		case value == "$ cd /":
			cur = root

		//set current directory to its parent
		case value == "$ cd ..":
			cur = cur.Parent

		// get directory name from line
		// call getDirectory method that finds the given dir from cur Children array
		case strings.HasPrefix(value, "$ cd "):
			lineContents := strings.Split(value, " ")
			cur = cur.getDirectory(lineContents[2])

			// handling the files or directories from ls command
		case value == "$ ls":
			for {
				//if end of input or a new command on next line break from inner for loop
				if i == len(lines)-1 || lines[i+1][0] == '$' {
					break
				}

				//otherwise increment index(move to next line)
				i++
				val := lines[i]
				lineContents2 := strings.Split(val, " ")

				//if not a directory add file to Children Slice
				if lineContents2[0] != "dir" {
					//fmt.Println(lineContents2[0])
					size, err := strconv.Atoi(lineContents2[0])
					if err != nil {
						panic(err)
					}
					cur.addFile(lineContents2[1], size)
					continue
				}

				//otherwise add the directory to Children slice of current Tree node
				cur.addDirectory(lineContents2[1])
			}
		}
	}

	total := findDirSize(root)
	fmt.Println(total)
}

func findDirSize(dir *Tree) int {
	total := 0

	if dir.Size <= 100000 {
		total += dir.Size
	}

	for _, v := range dir.Children {
		if v.Dir {
			total += findDirSize(v)
		}
	}

	return total
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
