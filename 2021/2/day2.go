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

func calculateDistance(inputData []string) int {
	horizontal := 0
	depth := 0
	for _, line := range inputData {
		if len(line) == 0 {
			continue
		}
		directions := strings.Split(line, " ")
		if len(directions) == 0 {
			continue
		}
		direction := directions[0]
		distance, _ := strconv.Atoi(directions[1])
		switch direction {
		case "forward":
			horizontal += distance
		case "up":
			depth -= distance
		case "down":
			depth += distance
		}
	}

	return horizontal * depth
}

func calculateAimedDistance(inputData []string) int {
	horizontal := 0
	aim := 0
	depth := 0
	for _, line := range inputData {
		if len(line) == 0 {
			continue
		}
		directions := strings.Split(line, " ")
		if len(directions) == 0 {
			continue
		}
		direction := directions[0]
		distance, _ := strconv.Atoi(directions[1])
		switch direction {
		case "forward":
			horizontal += distance
			depth += aim * distance
		case "up":
			aim -= distance
		case "down":
			aim += distance
		}
	}
	return horizontal * depth
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
		result = calculateDistance(inputData)
	} else {
		result = calculateAimedDistance(inputData)
	}

	fmt.Println(result)
}
