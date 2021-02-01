package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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

func validBirthYear(year int) bool {
	if year >= 1920 && year <= 2002 {
		return true
	}
	return false
}

func validIssueYear(year int) bool {
	if year >= 2010 && year <= 2020 {
		return true
	}
	return false
}

func validExpirationYear(year int) bool {
	if year >= 2020 && year <= 2030 {
		return true
	}
	return false
}

func validHeight(height string) bool {
	if len(height) < 4 {
		return false
	}
	value, err := strconv.Atoi(height[:len(height)-2])
	if err != nil {
		// fmt.Fprintf(os.Stderr, "Failed to parse height: %s\n", err)
		return false
	}
	unit := height[len(height)-2:]
	if unit == "cm" {
		if value >= 150 && value <= 193 {
			return true
		}
	}
	if unit == "in" {
		if value >= 59 && value <= 76 {
			return true
		}
	}
	return false
}

func validHairColor(color string) bool {
	valid := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	if valid.MatchString(color) {
		return true
	}
	return false
}

func validEyeColor(color string) bool {
	valid := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	if valid.MatchString(color) {
		return true
	}
	return false
}

func validPassportNumber(passport string) bool {
	valid := regexp.MustCompile(`^[0-9]{9}$`)
	if valid.MatchString(passport) {
		return true
	}
	return false
}

func checkValidPassport(passport *Passport, strict bool) {
	if strict {
		if validBirthYear(passport.birthYear) &&
			validIssueYear(passport.issueYear) &&
			validExpirationYear(passport.expireYear) &&
			validHeight(passport.height) &&
			validHairColor(passport.hairColor) &&
			validEyeColor(passport.eyeColor) &&
			validPassportNumber(passport.passportID) {
			passport.valid = true
		}
	} else {
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

}

func parsePassport(passportFields []string, strict bool) Passport {
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
	checkValidPassport(&passport, strict)
	return passport
}

func parseInputData(inputData []string, strict bool) []Passport {
	var passportFields []string
	var passports []Passport

	for _, line := range inputData {
		if line != "" {
			passportFields = append(passportFields, line)
		} else {
			passports = append(passports, parsePassport(passportFields, strict))
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

	flag.StringVar(&fileName, "f", "input", "Input file")
	flag.BoolVar(&part2, "2", false, "Compute part 2 of the exercise")
	flag.Parse()

	inputData, err := readFile(fileName)
	if err != nil {
		die(err)
	}

	passports = parseInputData(inputData, part2)
	result = countValidPassports(passports)

	fmt.Println(result)
}
