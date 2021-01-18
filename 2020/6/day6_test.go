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

func TestFindAnswered(t *testing.T) {
	expected := map[string]bool{
		"a": true,
		"b": true,
	}

	response := findAnswered([]string{"a", "b"})
	if !reflect.DeepEqual(expected, response) {
		t.Log("Error, failed to find the answered questions", response)
		t.Fail()
	}

	expected = map[string]bool{}
	response = findAnswered([]string{})
	if !reflect.DeepEqual(expected, response) {
		t.Log("Error, failed to find the answered questions", response)
		t.Fail()
	}
}

func TestCollectForms(t *testing.T) {
	response := collectForms([]string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	})
	if !reflect.DeepEqual(response, map[int][]string{
		0: {"abc"},
		1: {"a", "b", "c"},
		2: {"ab", "ac"},
		3: {"a", "a", "a", "a"},
		4: {"b"},
	}) {
		t.Log("Error, got unexpected form collection", response)
		t.Fail()
	}

	response = collectForms([]string{
		"",
		"",
	})
	if !reflect.DeepEqual(response, map[int][]string{}) {
		t.Log("Error, expected empty form collection, got", response)
		t.Fail()
	}

	response = collectForms([]string{})
	if !reflect.DeepEqual(response, map[int][]string{}) {
		t.Log("Error, expected empty form collection, got", response)
		t.Fail()
	}
}

func TestCountAnswerd(t *testing.T) {
	var response int
	response = countAnswered(map[int][]string{
		0: {"abc"},
		1: {"a", "b", "c"},
		2: {"ab", "ac"},
		3: {"a", "a", "a", "a"},
		4: {"b"},
	})
	if response != 11 {
		t.Log("Expected 11, got", response)
		t.Fail()
	}
}

func TestCountEveryoneAnswerd(t *testing.T) {
	var response int
	response = countEveryoneAnswered(map[int][]string{
		0: {"abc"},
		1: {"a", "b", "c"},
		2: {"ab", "ac"},
		3: {"a", "a", "a", "a"},
		4: {"b"},
	})
	if response != 6 {
		t.Log("Expected 6, got", response)
		t.Fail()
	}
}
