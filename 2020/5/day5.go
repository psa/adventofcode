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

func findSeatRow(seat string) int {
	var seatNumber []byte
	if len(seat) != 7 {
		fmt.Fprintf(os.Stderr, "Bad seat ID: %s, length %d instead of 7\n", seat, len(seat))
		return 0
	}
	for _, c := range seat {
		switch c {
		case 'F':
			seatNumber = append(seatNumber, '0')
		case 'B':
			seatNumber = append(seatNumber, '1')
		default:
			fmt.Fprintf(os.Stderr, "Error, bad position identifier in %s: %c\n", seat, c)
			return 0
		}
	}
	if row, err := strconv.ParseUint(string(seatNumber), 2, 8); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return 0
	} else {
		return int(row)
	}
}

func findSeatColumn(seat string) int {
	var seatNumber []byte
	if len(seat) != 3 {
		fmt.Printf("Bad seat ID: %s, length %d instead of 3\n", seat, len(seat))
		return 0
	}
	for _, c := range seat {
		switch c {
		case 'L':
			seatNumber = append(seatNumber, '0')
		case 'R':
			seatNumber = append(seatNumber, '1')
		default:
			fmt.Printf("Error, bad position identifier in %s: %c\n", seat, c)
			return 0
		}
	}
	if row, err := strconv.ParseUint(string(seatNumber), 2, 4); err != nil {
		fmt.Println(err)
		return 0
	} else {
		return int(row)
	}
}

func generateSeatID(row int, column int) int {
	if row < 0 || row > 127 {
		fmt.Println("Row is out of bounds", row)
		return 0
	}
	if column < 0 || column > 7 {
		fmt.Println("Column is out of bounds", column)
		return 0
	}
	return (row * 8) + column
}

func generateSeatIDs(inputData []string) []int {
	var seatIDs []int

	for _, line := range inputData {
		if len(line) == 0 {
			continue
		}
		if len(line) != 10 {
			fmt.Fprintf(os.Stderr, "Error, short line: %v\n", line)
			continue
		}
		row := line[:7]
		column := line[7:]
		seatIDs = append(seatIDs, generateSeatID(findSeatRow(row), findSeatColumn(column)))
	}
	return seatIDs
}

func findHighestSeatID(seatIDs []int) (int, error) {
	if len(seatIDs) < 1 {
		return 0, errors.New("Not passed any seat IDs")
	}
	sort.Ints(seatIDs)
	highest := seatIDs[len(seatIDs)-1:]
	return highest[0], nil
}

func findMissingSeat(seatIDs []int) (int, error) {
	var last int
	if 0 == len(seatIDs) {
		return 0, errors.New("Not passed any seat IDs")
	}
	for _, seat := range seatIDs {
		if 0 == last {
			last = seat
			continue
		}
		if last != seat-1 {
			return seat - 1, nil
		}
		last = seat
	}
	return 0, errors.New("No missing seat found")
}

func main() {
	var fileName string
	var part2 bool

	flag.StringVar(&fileName, "f", "input", "Input file")
	flag.BoolVar(&part2, "2", false, "Compute part 2 of the exercise")
	flag.Parse()

	inputData, err := readFile(fileName)
	if err != nil {
		die(err)
	}

	seatIDs := generateSeatIDs(inputData)
	if part2 {
		sort.Ints(seatIDs)
		seat, err := findMissingSeat(seatIDs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to find a seat")
			os.Exit(1)
		}
		fmt.Println(seat)
	} else {
		result, err := findHighestSeatID(seatIDs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		fmt.Println(result)
	}
}
