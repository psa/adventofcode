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

func TestCheckValidPassport(t *testing.T) {
	passport := Passport{
		birthYear:  2000,
		issueYear:  2000,
		expireYear: 2000,
		height:     "200cm",
		hairColor:  "#fffff",
		eyeColor:   "grn",
		passportID: "019123",
		countryID:  0,
		valid:      false,
	}
	checkValidPassport(&passport)
	if passport.valid != true {
		t.Log("Valid passport marked invalid:", passport)
		t.Fail()
	}
}

func TestCheckInvalidPassport(t *testing.T) {
	passport := Passport{
		birthYear:  2000,
		issueYear:  2000,
		expireYear: 0,
		height:     "200cm",
		hairColor:  "#fffff",
		eyeColor:   "grn",
		passportID: "019123",
		countryID:  0,
		valid:      false,
	}
	checkValidPassport(&passport)
	if passport.valid == true {
		t.Log("Invalid passport marked valid:", passport)
		t.Fail()
	}
}

func TestCountValidPassports(t *testing.T) {
	passports := []Passport{
		Passport{
			birthYear:  2000,
			issueYear:  2000,
			expireYear: 0,
			height:     "200cm",
			hairColor:  "#fffff",
			eyeColor:   "grn",
			passportID: "019123",
			countryID:  0,
			valid:      false,
		},
		Passport{
			birthYear:  2000,
			issueYear:  2000,
			expireYear: 2020,
			height:     "200cm",
			hairColor:  "#fffff",
			eyeColor:   "grn",
			passportID: "019123",
			countryID:  0,
			valid:      true,
		},
	}
	count := countValidPassports(passports)
	if count != 1 {
		t.Log("Expected 1 valid passport, got:", count)
		t.Fail()
	}
}

func TestParsePassport(t *testing.T) {
	passport := Passport{
		birthYear:  1937,
		issueYear:  2017,
		expireYear: 2020,
		height:     "183cm",
		hairColor:  "#fffffd",
		eyeColor:   "gry",
		passportID: "860033327",
		countryID:  147,
		valid:      true,
	}
	passportFields := []string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
	}
	result := parsePassport(passportFields)
	if !reflect.DeepEqual(passport, result) {
		t.Log("Error, expected and actual passport differ", passport, result)
		t.Fail()
	}
}
