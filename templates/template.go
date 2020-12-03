package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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

func parseInputData(inputData []string) int {
	return 0
}

func main() {
	var fileName string
	var part2 bool
	var result int

	flag.StringVar(&fileName, "f", "input/XXX", "Input file")
	flag.BoolVar(&part2, "2", false, "Compute part 2 of the exercise")
	flag.Parse()

	inputData, err := readFile(fileName)
	if err != nil {
		die(err)
	}

	result = parseInputData(inputData)

	fmt.Println(result)
}
