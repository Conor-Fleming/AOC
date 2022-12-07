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

	for i := 0; i < len(lines); i++ {
		value := lines[i]
		if value == "$ cd /" {
			cur = root
			continue
		}
		if value == "$ cd .." {
			cur = cur.Parent
			continue
		}
		if strings.HasPrefix(value, "$ cd ") {
			lineContents := strings.Split(value, " ")
			cur = cur.getDirectory(lineContents[2])
			continue
		}
		if value == "$ ls" {
			for {
				if i == len(lines)-1 || lines[i+1][0] == '$' {
					break
				}

				i++
				val := lines[i]
				lineContents2 := strings.Split(val, " ")

				if lineContents2[0] != "dir" {
					//fmt.Println(lineContents2[0])
					size, err := strconv.Atoi(lineContents2[0])
					if err != nil {
						panic(err)
					}
					cur.addFile(lineContents2[1], size)
					continue
				}

				cur.addDirectory(lineContents2[1])
			}
			continue
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
