package main

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

var test_data = []string{
	"1",
	"2",
	"3",
	"",
	"100",
	"200",
	"300",
	"",
	"10",
	"20",
	"30",
}

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

func TestCalculateLoads(t *testing.T) {
	expected := []int{6, 600, 60}
	result, err := calculateLoads(test_data)
	if nil != err {
		t.Log("Unexpected error:", err)
		t.Fail()
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Log("Expected:", expected, "got:", result)
			t.Fail()
		}
	}
}

func TestCalculateLoadsErrors(t *testing.T) {
	data := []string{"A", "B"}
	var result []int
	var err error
	result, err = calculateLoads(data)
	if nil == err {
		t.Log("Expected error, did not get one. Result:", result)
		t.Fail()
	}

	data = []string{"", "", ""}
	result, err = calculateLoads(data)
	if nil == err {
		t.Log("Expected error, did not get one. Result:", result)
		t.Fail()
	}
}
