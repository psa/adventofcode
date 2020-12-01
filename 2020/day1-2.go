package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var expenses []int
	var expensesSubset []int
	var expensesSubsetSubset []int // what a crappy hack
	var loop int
	file, err := os.Open("input/1")
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
}
