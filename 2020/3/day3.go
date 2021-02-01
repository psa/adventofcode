package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func readTrees(filename string) []string {
	var trees []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trees = append(trees, scanner.Text())
	}
	return trees
}

func parseTreeLines(treeLines []string) (int, map[int][]int) {
	var trees = make(map[int][]int)
	var counter int
	var length int

	for _, line := range treeLines {
		length = len(line)
		for pos, item := range line {
			if item == '#' {
				trees[counter] = append(trees[counter], pos)
			}
		}
		counter++
	}
	return length, trees
}

func hitTree(pos int, trees []int) bool {
	for _, tree := range trees {
		if pos == tree {
			return true
		}
	}
	return false
}

func extendTrees(length int, trees []int, position int) []int {
	var newTrees []int
	loop := 1
	for _, i := range trees {
		newTrees = append(newTrees, i)
	}
	for newTrees[len(newTrees)-1] < position {
		for _, tree := range trees {
			newTrees = append(newTrees, tree+(length*loop))
		}
		loop++
	}
	return newTrees
}

func scanTrees(length int, trees map[int][]int, right int, down int) int {
	var position int
	var treeHits int
	//for line, _ := range trees {
	for line := 0; line < len(trees); line++ {
		treeLine := trees[line]
		if line%down != 0 {
			continue
		}
		if position >= treeLine[len(treeLine)-1] {
			treeLine = extendTrees(length, treeLine, position)
		}
		if hitTree(position, treeLine) {
			treeHits++
		}
		position += right
	}
	return treeHits
}

func main() {
	var treeLines []string
	var fileName string
	var result int
	var down int
	var right int
	var trees = make(map[int][]int)
	var length int
	var part2 bool

	flag.StringVar(&fileName, "f", "input", "Input file")
	flag.IntVar(&down, "d", 1, "Points to travel down")
	flag.IntVar(&right, "r", 3, "Points to travel right")
	flag.BoolVar(&part2, "2", false, "Compute part 2 of the exercise")
	flag.Parse()

	treeLines = readTrees(fileName)
	if nil == treeLines {
		fmt.Println("Input file was empty")
		os.Exit(1)
	}

	length, trees = parseTreeLines(treeLines)
	if trees == nil {
		os.Exit(1)
	}

	if part2 {
		result = 1
		result *= scanTrees(length, trees, 1, 1)
		result *= scanTrees(length, trees, 3, 1)
		result *= scanTrees(length, trees, 5, 1)
		result *= scanTrees(length, trees, 7, 1)
		result *= scanTrees(length, trees, 1, 2)
	} else {
		result = scanTrees(length, trees, right, down)
	}

	fmt.Println(result)
}
