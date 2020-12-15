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

func countAnswered(forms map[int][]string) int {
	var count int
	for _, form := range forms {
		answeredQuestions := findAnswered(form)
		count += len(answeredQuestions)
	}
	return count
}

func findEveryoneAnswered(form []string) map[string]bool {
	questions := make(map[string]int)
	everyoneAnswered := make(map[string]bool)
	groupSize := len(form)
	for _, line := range form {
		for _, answer := range strings.Split(line, "") {
			if answer != "" {
				questions[answer]++
			}
		}
	}

	for question, count := range questions {
		if count == groupSize {
			everyoneAnswered[question] = true
		}
	}
	return everyoneAnswered
}

func countEveryoneAnswered(forms map[int][]string) int {
	var count int
	for _, form := range forms {
		answeredQuestions := findEveryoneAnswered(form)
		count += len(answeredQuestions)
	}
	return count
}

func collectForms(inputData []string) map[int][]string {
	customsForms := make(map[int][]string)
	var count int

	for _, line := range inputData {
		if line != "" {
			customsForms[count] = append(customsForms[count], line)
		} else {
			count++
		}
	}
	return customsForms
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

	forms := collectForms(inputData)
	if part2 {
		result = countEveryoneAnswered(forms)
	} else {
		result = countAnswered(forms)
	}

	fmt.Println(result)
}
