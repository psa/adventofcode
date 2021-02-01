package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passwordData struct {
	Min       int
	Max       int
	Character string
	Password  string
}

func readPasswords(filename string) []string {
	var passwords []string
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passwords = append(passwords, scanner.Text())
	}
	return passwords
}

func parsePasswordLines(passwordLines []string) []passwordData {
	var passwords []passwordData

	// 1-10 j: vrfjljjwbsv
	re := regexp.MustCompile(`^(?P<min>\d+)-(?P<max>\d+) (?P<char>[a-zA-Z]): (?P<password>\w+)$`)
	for _, line := range passwordLines {
		if !re.MatchString(line) {
			fmt.Println("Bad line: ", line)
		} else {
			subMatch := re.FindStringSubmatch(line)
			min, err := strconv.Atoi(subMatch[1])
			if err != nil {
				fmt.Println("Unable to convert text to int", err)
				return nil
			}
			max, err := strconv.Atoi(subMatch[2])
			if err != nil {
				fmt.Println("Unable to convert text to int", err)
				return nil
			}
			passwords = append(passwords, passwordData{
				min,
				max,
				subMatch[3],
				subMatch[4],
			})
		}
	}
	return passwords
}

func scanPasswords(passwords []passwordData) int {
	var correct int
	for _, entry := range passwords {
		count := strings.Count(entry.Password, entry.Character)
		if count >= entry.Min && count <= entry.Max {
			correct++
		}
	}
	return correct
}

func scanPasswordsNewPolicy(passwords []passwordData) int {
	var correct int
	var first bool
	var second bool
	for _, entry := range passwords {
		// Re-base to zero from one so it lines up with arrays being zero based
		min := entry.Min - 1
		max := entry.Max - 1
		if min >= len(entry.Password) {
			first = false
		} else {
			first = strings.Contains(string(entry.Password[min]), entry.Character)
		}
		if max >= len(entry.Password) {
			second = false
		} else {
			second = strings.Contains(string(entry.Password[max]), entry.Character)
		}
		if (first || second) && !(first && second) {
			correct++
		}
	}
	return correct
}

func main() {
	var passwordLines []string
	var fileName string
	var passwords []passwordData
	var newPolicy bool
	var result int

	flag.StringVar(&fileName, "f", "input", "Input file")
	flag.BoolVar(&newPolicy, "n", false, "Use New Policy")
	flag.Parse()

	passwordLines = readPasswords(fileName)
	if nil == passwordLines {
		fmt.Println("Input file was empty")
		os.Exit(1)
	}

	passwords = parsePasswordLines(passwordLines)
	if passwords == nil {
		os.Exit(1)
	}

	if newPolicy {
		result = scanPasswordsNewPolicy(passwords)
	} else {
		result = scanPasswords(passwords)
	}
	fmt.Println(result)
}
