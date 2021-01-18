package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var fileName string
	var part2 bool
	var expenses []int
	var expensesSubset []int
	var expensesSubsetSubset []int // what a crappy hack
	var loop int

	flag.StringVar(&fileName, "f", "input", "Input file")
	flag.BoolVar(&part2, "2", false, "Compute part 2 of the exercise")
	flag.Parse()

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Unable to convert text to int", err)
		}
		expenses = append(expenses, data)
	}

	sort.Ints(expenses)

	if part2 {
		for i := 0; i < len(expenses); i++ {
			loop++
			expensesSubset = expenses[loop:]
			expensesSubsetSubset = expenses[loop+1:]
			for j := 0; j < len(expensesSubset); j++ {
				for k := 0; k < len(expensesSubsetSubset); k++ {
					if 2020 == expenses[i]+expensesSubset[j]+expensesSubsetSubset[k] {
						fmt.Println(expenses[i] * expensesSubset[j] * expensesSubsetSubset[k])
						return
					}
					if 2020 < expenses[i]+expensesSubset[j]+expensesSubsetSubset[k] {
						break
					}
				}
			}
		}
	} else {
		for i := 0; i < len(expenses); i++ {
			loop++
			expensesSubset = expenses[loop:]
			for j := 0; j < len(expensesSubset); j++ {
				if 2020 == expenses[i]+expensesSubset[j] {
					fmt.Println(expenses[i] * expensesSubset[j])
					return
				}
			}
		}
	}

}
