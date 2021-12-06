package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/* Errors should be printed to STDERR not STDOUT */
func die(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}

func readFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	/* Split on an empty returns an array of length 1 */
	if len(data) == 0 || string(data) == "\n" || len(lines) == 0 {
		return nil, errors.New("Empty file")
	}

	return lines, nil
}

func comparePrevious(inputData []string) int {
	var count = -1
	var previous = 0
	for _, line := range inputData {
		num, _ := strconv.Atoi(line)
		if num > previous {
			count++
		}
		previous = num
	}
	return count
}

func comparePreviousThree(inputData []string) int {
	var count = 0
	var sum = 0
	var lastSum = 0
	nums := []int{0, 0, 0}
	for idx, line := range inputData {
		num, _ := strconv.Atoi(line)
		nums = nums[1:]
		nums = append(nums, num)
		sum = nums[0] + nums[1] + nums[2]
		if idx < 3 {
			lastSum = sum
			continue
		}
		if sum > lastSum {
			count++
		}
		lastSum = sum
	}
	return count
}

func main() {
	var fileName string
	var part2 bool
	var result int

	flag.StringVar(&fileName, "f", "input", "Input file")
	flag.BoolVar(&part2, "2", false, "Compute part 2 of the exercise")
	flag.Parse()

	inputData, err := readFile(fileName)
	if err != nil {
		die(err)
	}

	if !part2 {
		result = comparePrevious(inputData)
	} else {
		result = comparePreviousThree(inputData)
	}

	fmt.Println(result)
}
