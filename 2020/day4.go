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

type Passport struct {
	birthYear  int    // byr
	issueYear  int    // iyr
	expireYear int    // eyr
	height     string // hgt
	hairColor  string // hcl
	eyeColor   string // ecl
	passportID string // pid, can have a leading zero, so int is a bad idea
	countryID  int    // cid, optional in part 1
	valid      bool
}

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

func checkValidPassport(passport *Passport) {
	if passport.birthYear != 0 &&
		passport.issueYear != 0 &&
		passport.expireYear != 0 &&
		passport.height != "" &&
		passport.hairColor != "" &&
		passport.eyeColor != "" &&
		passport.passportID != "" {
		passport.valid = true
	}
}

func parsePassport(passportFields []string) Passport {
	var passport Passport
	for _, line := range passportFields {
		fields := strings.Split(line, " ")
		for _, field := range fields {
			fieldParts := strings.Split(field, ":")
			switch fieldParts[0] {
			case "byr":
				year, err := strconv.Atoi(fieldParts[1])
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to parse int: %s\n", err)
				}
				passport.birthYear = year
			case "iyr":
				year, err := strconv.Atoi(fieldParts[1])
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to parse int: %s\n", err)
				}
				passport.issueYear = year
			case "eyr":
				year, err := strconv.Atoi(fieldParts[1])
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to parse int: %s\n", err)
				}
				passport.expireYear = year
			case "hgt":
				passport.height = fieldParts[1]
			case "hcl":
				passport.hairColor = fieldParts[1]
			case "ecl":
				passport.eyeColor = fieldParts[1]
			case "pid":
				passport.passportID = fieldParts[1]
			case "cid":
				country, err := strconv.Atoi(fieldParts[1])
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to parse int: %s\n", err)
				}
				passport.countryID = country
			}
		}
	}
	checkValidPassport(&passport)
	return passport
}

func parseInputData(inputData []string) []Passport {
	var passportFields []string
	var passports []Passport

	for _, line := range inputData {
		if line != "" {
			passportFields = append(passportFields, line)
		} else {
			passports = append(passports, parsePassport(passportFields))
			passportFields = []string{}
		}
	}
	return passports
}

func countValidPassports(passports []Passport) int {
	var counter int
	for _, passport := range passports {
		if passport.valid {
			counter++
		}
	}
	return counter
}

func main() {
	var fileName string
	var part2 bool
	var result int
	var passports []Passport

	flag.StringVar(&fileName, "f", "input/4", "Input file")
	flag.BoolVar(&part2, "2", false, "Compute part 2 of the exercise")
	flag.Parse()

	inputData, err := readFile(fileName)
	if err != nil {
		die(err)
	}

	passports = parseInputData(inputData)
	result = countValidPassports(passports)

	fmt.Println(result)
}
