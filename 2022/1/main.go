package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
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

func calculateLoads(food []string) ([]int, error) {
	var loads []int
	var currentLoad int
	hasData := false
	for _, item := range food {
		if len(item) > 0 {
			hasData = true
			load, err := strconv.Atoi(item)
			if nil != err {
				return nil, err
			}
			currentLoad += load
		} else {
			loads = append(loads, currentLoad)
			currentLoad = 0
		}
	}
	if len(loads) < 1 || !hasData {
		return nil, errors.New("Unable to find any loads")
	}
	return loads, nil
}

func findHeaviestLoad(food []int) int {
	heaviest := 0
	for _, load := range food {
		if load > heaviest {
			heaviest = load
		}
	}
	return heaviest
}

func findTopThreeTotal(food []int) (int, error) {
	total := 0
	sort.Ints(food)
	if len(food) < 3 {
		return -1, errors.New("Not enough elves, need minimum of 3")
	}
	lastThree := food[len(food)-3:]
	for _, i := range lastThree {
		total += i
	}
	return total, nil
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

	loads, err := calculateLoads(inputData)
	if nil != err {
		die(err)
	}

	if !part2 {
		result = findHeaviestLoad(loads)
	} else {
		result, err = findTopThreeTotal(loads)
		if nil != err {
			die(err)
		}
	}
	fmt.Println(result)
}
