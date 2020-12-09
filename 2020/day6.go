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

func findAnswered(form []string) map[string]bool {
	questions := make(map[string]bool)
	for _, line := range form {
		for _, answer := range strings.Split(line, "") {
			if answer != "" {
				questions[answer] = true
			}
		}
	}
	return questions
}

func countAnswered(inputData []string) int {
	var customsForms []string
	var count int

	for _, line := range inputData {
		if line != "" {
			customsForms = append(customsForms, line)
		} else {
			answeredQuestions := findAnswered(customsForms)
			count += len(answeredQuestions)
			customsForms = []string{}
		}
	}
	// Clean up if no trailing newline, OK if there is as it's reset above on
	// completion of a form line
	answeredQuestions := findAnswered(customsForms)
	count += len(answeredQuestions)

	return count
}

func main() {
	var fileName string
	var part2 bool
	var result int

	flag.StringVar(&fileName, "f", "input/6", "Input file")
	flag.BoolVar(&part2, "2", false, "Compute part 2 of the exercise")
	flag.Parse()

	inputData, err := readFile(fileName)
	if err != nil {
		die(err)
	}

	result = countAnswered(inputData)

	fmt.Println(result)
}
