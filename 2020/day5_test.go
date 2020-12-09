package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func createTestFile(contents []byte, t *testing.T) string {
	tempFile, err := ioutil.TempFile("/tmp", "unit-test.*")
	if err != nil {
		t.Log("Failed to create temp file", err)
		t.Fail()
	}

	if len(contents) != 0 {
		if _, err := tempFile.Write(contents); err != nil {
			t.Log("Failed to write to file:", err)
			t.Fail()
		}
	}

	if err := tempFile.Close(); err != nil {
		t.Log("Failed to write to file:", err)
		t.Fail()
	}

	return tempFile.Name()
}

func TestCreateTestFile(t *testing.T) {
	contents := []byte("Line1\nLine2")
	fileName := createTestFile(contents, t)
	defer os.Remove(fileName)

	readContents, err := readFile(fileName)
	if err != nil {
		t.Log("Failed to read file:", err)
		t.Fail()
	}

	if !reflect.DeepEqual(strings.Split(string(contents), "\n"), readContents) {
		t.Log("Error, contents of test file differs", readContents)
		t.Fail()
	}
}

func TestReadFile(t *testing.T) {
	contents := []byte("Line1\nLine2")
	fileName := createTestFile(contents, t)
	defer os.Remove(fileName)

	readContents, err := readFile(fileName)
	if err != nil {
		t.Log("Failed to read file:", err)
		t.Fail()
	}

	if !reflect.DeepEqual(strings.Split(string(contents), "\n"), readContents) {
		t.Log("Error, contents of test file differs", readContents)
		t.Fail()
	}
}

func TestReadFileSingleLine(t *testing.T) {
	contents := []byte("Line1")
	fileName := createTestFile(contents, t)
	defer os.Remove(fileName)

	readContents, err := readFile(fileName)
	if err != nil {
		t.Log("Failed to read file:", err)
		t.Fail()
	}

	if !reflect.DeepEqual(strings.Split(string(contents), "\n"), readContents) {
		t.Log("Error, contents of test file differs", readContents)
		t.Fail()
	}
}

func TestReadFileNonexistent(t *testing.T) {
	_, err := readFile("unit-test.does-not-exist")
	if err == nil {
		t.Log("Failed to error on non-existent file")
		t.Fail()
	}
}

func TestReadFileEmpty(t *testing.T) {
	tempFile, err := ioutil.TempFile("/tmp", "unit-test.*")
	if err != nil {
		t.Log("Failed to create temp file", err)
		t.Fail()
	}
	defer os.Remove(tempFile.Name())

	if err := tempFile.Close(); err != nil {
		t.Log("Failed to write to file:", err)
		t.Fail()
	}

	readContents, err := readFile(tempFile.Name())
	if err == nil {
		t.Log("Should error on empty file, got:", readContents)
		t.Fail()
	}
}

func TestReadFileSingleNewline(t *testing.T) {
	contents := []byte("\n")
	fileName := createTestFile(contents, t)
	defer os.Remove(fileName)

	readContents, err := readFile(fileName)
	if err == nil {
		t.Log("Got contents where there should be none:", readContents)
		t.Fail()
	}
}

func TestFindSeatRow(t *testing.T) {
	seat := "FBFBBFF"
	row := findSeatRow(seat)
	if row != 44 {
		t.Log("Expected row 44, got", row)
		t.Fail()
	}

	seat = "FBFBBF"
	row = findSeatRow(seat)
	if row != 0 {
		t.Log("Got row 0 (too short), got", row)
		t.Fail()
	}

	seat = "FBFBBFFF"
	row = findSeatRow(seat)
	if row != 0 {
		t.Log("Got row 0 (too long), got", row)
		t.Fail()
	}

	seat = "9FBFBBF"
	row = findSeatRow(seat)
	if row != 0 {
		t.Log("Got row 0 (non-F|B character), got", row)
		t.Fail()
	}
}

func TestFindSeatColumn(t *testing.T) {
	column := findSeatColumn("RLR")
	if column != 5 {
		t.Log("Expected column 5, got", column)
		t.Fail()
	}

	column = findSeatColumn("RR")
	if column != 0 {
		t.Log("Got column 0 (too short), got", column)
		t.Fail()
	}

	column = findSeatColumn("RLRL")
	if column != 0 {
		t.Log("Got column 0 (too long), got", column)
		t.Fail()
	}

	column = findSeatColumn("9RL")
	if column != 0 {
		t.Log("Got column 0 (non-L|R character), got", column)
		t.Fail()
	}
}

func TestGenerateSeatID(t *testing.T) {
	seatID := generateSeatID(44, 5)
	if seatID != 357 {
		t.Log("Wrong seat ID, expected 357, got", seatID)
		t.Fail()
	}

	// Top of allowable row
	seatID = generateSeatID(127, 1)
	if seatID != 1017 {
		t.Log("Wrong seat ID, expected 1017, got", seatID)
		t.Fail()
	}

	// Top of allowable column
	seatID = generateSeatID(1, 1)
	if seatID != 9 {
		t.Log("Wrong seat ID, expected 9, got", seatID)
		t.Fail()
	}

	seatID = generateSeatID(-1, 5)
	if seatID != 0 {
		t.Log("Wrong seat ID, expected 0 (negative row), got", seatID)
		t.Fail()
	}

	seatID = generateSeatID(128, 1)
	if seatID != 0 {
		t.Log("Wrong seat ID, expected 0 (high row), got", seatID)
		t.Fail()
	}

	seatID = generateSeatID(1, -1)
	if seatID != 0 {
		t.Log("Wrong seat ID, expected 0 (negative column), got", seatID)
		t.Fail()
	}

	seatID = generateSeatID(1, 8)
	if seatID != 0 {
		t.Log("Wrong seat ID, expected 0 (high column), got", seatID)
		t.Fail()
	}
}

func TestFindHighestSeatID(t *testing.T) {
	var seatID int
	var err error
	seatID, err = findHighestSeatID([]int{4, 6, 7, 2, 1, 9})
	if seatID != 9 {
		t.Log("Wrong seat ID, expected 9, got", seatID)
		t.Fail()
	}

	seatID, err = findHighestSeatID([]int{})
	if err == nil {
		t.Log("Expected an error, for an empty seat array but didn't get one")
		t.Fail()
	}
}

func TestFindMissingSeat(t *testing.T) {
	var seat int
	var err error

	seat, err = findMissingSeat([]int{1, 2, 3, 5, 6, 7})
	if err != nil {
		t.Log("Unexpectedly got an error testing findMissingSeat", err)
		t.Fail()
	}
	if seat != 4 {
		t.Log("Expected seat 4, got", seat)
		t.Fail()
	}

	seat, err = findMissingSeat([]int{1, 2})
	if err == nil {
		t.Log("Expected an error and got none testing no missing seat")
		t.Fail()
	}

	seat, err = findMissingSeat([]int{})
	if err == nil {
		t.Log("Expected an error and testing empty input for findMissingSeat")
		t.Fail()
	}

}
