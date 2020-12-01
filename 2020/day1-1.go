package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var expenses []int
	var expensesSubset []int
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
