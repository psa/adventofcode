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
