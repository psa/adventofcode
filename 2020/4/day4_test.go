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
	checkValidPassport(&passport, false)
	if passport.valid != true {
		t.Log("Valid passport marked invalid:", passport)
		t.Fail()
	}

	passport = Passport{
		birthYear:  2000,
		issueYear:  2010,
		expireYear: 2020,
		height:     "182cm",
		hairColor:  "#123456",
		eyeColor:   "grn",
		passportID: "012345678",
		countryID:  0,
		valid:      false,
	}
	checkValidPassport(&passport, true)
	if passport.valid != true {
		t.Log("Valid passport marked invalid with strict conditions:", passport)
		t.Fail()
	}

	passport = Passport{
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
	checkValidPassport(&passport, false)
	if passport.valid == true {
		t.Log("Invalid passport marked valid:", passport)
		t.Fail()
	}

	passport = Passport{
		birthYear:  1000,
		issueYear:  2010,
		expireYear: 2020,
		height:     "182cm",
		hairColor:  "#123456",
		eyeColor:   "grn",
		passportID: "012345678",
		countryID:  0,
		valid:      false,
	}
	checkValidPassport(&passport, true)
	if passport.valid == true {
		t.Log("Invalid passport marked valid with strict conditions:", passport)
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
	result := parsePassport(passportFields, false)
	if !reflect.DeepEqual(passport, result) {
		t.Log("Error, expected and actual passport differ", passport, result)
		t.Fail()
	}
}

func TestValidBirthYear(t *testing.T) {
	if validBirthYear(1919) {
		t.Log("Error, 1919 considered a valid birth year")
		t.Fail()
	}
	if !validBirthYear(1920) {
		t.Log("Error, 1920 not considered a valid birth year")
		t.Fail()
	}
	if !validBirthYear(2002) {
		t.Log("Error, 2002 not considered a valid birth year")
		t.Fail()
	}
	if validBirthYear(2003) {
		t.Log("Error, 2003 considered a valid birth year")
		t.Fail()
	}
}

func TestValidIssueYear(t *testing.T) {
	if validIssueYear(2009) {
		t.Log("Error, 2009 considered a valid issue year")
		t.Fail()
	}
	if !validIssueYear(2010) {
		t.Log("Error, 2010 not considered a valid issue year")
		t.Fail()
	}
	if !validIssueYear(2020) {
		t.Log("Error, 2020 not considered a valid issue year")
		t.Fail()
	}
	if validIssueYear(2021) {
		t.Log("Error, 2021 considered a valid issue year")
		t.Fail()
	}
}

func TestValidExpirationYear(t *testing.T) {
	if validExpirationYear(2019) {
		t.Log("Error, 2019 considered a valid expiration year")
		t.Fail()
	}
	if !validExpirationYear(2020) {
		t.Log("Error, 2020 not considered a valid expiration year")
		t.Fail()
	}
	if !validExpirationYear(2030) {
		t.Log("Error, 2030 not considered a valid expiration year")
		t.Fail()
	}
	if validExpirationYear(2031) {
		t.Log("Error, 2031 considered a valid expiration year")
		t.Fail()
	}
}

func TestValidHeight(t *testing.T) {
	// Metric
	if validHeight("149cm") {
		t.Log("Error, 149cm considered a valid height")
		t.Fail()
	}
	if !validHeight("150cm") {
		t.Log("Error, 150cm not considered a valid height")
		t.Fail()
	}
	if !validHeight("193cm") {
		t.Log("Error, 193cm not considered a valid height")
		t.Fail()
	}
	if validHeight("194cm") {
		t.Log("Error, 194cm considered a valid height")
		t.Fail()
	}

	// When will we stop wasting our lives catering to this?
	if validHeight("58in") {
		t.Log("Error, 58in considered a valid height")
		t.Fail()
	}
	if !validHeight("59in") {
		t.Log("Error, 59in not considered a valid height")
		t.Fail()
	}
	if !validHeight("76in") {
		t.Log("Error, 76in not considered a valid height")
		t.Fail()
	}
	if validHeight("77in") {
		t.Log("Error, 77in considered a valid height")
		t.Fail()
	}

	// Random string
	if validHeight("invalid") {
		t.Log("Error, invalid considered a valid height")
	}

	// Missing digits
	if validHeight("cm") {
		t.Log("Error, cm considered a valid height")
	}

	// Invalid unit
	if validHeight("180xx") {
		t.Log("Error, 180xx considered a valid height")
	}

	// Wrong order
	if validHeight("cm150") {
		t.Log("Error, cm150 considered a valid height")
	}

	// Too short
	if validHeight("x") {
		t.Log("Error, x considered a valid height")
	}

}

func TestValidHairColor(t *testing.T) {
	if !validHairColor("#0099af") {
		t.Log("Error, #0099af not considered a valid hair color")
		t.Fail()
	}
	// too long
	if validHairColor("#0099afa") {
		t.Log("Error, #0099afa considered a valid hair color")
		t.Fail()
	}
	// too short
	if validHairColor("#0099a") {
		t.Log("Error, #0099a considered a valid hair color")
		t.Fail()
	}
	// Hex out of bounds
	if validHairColor("#0099ag") {
		t.Log("Error, #0099ag considered a valid hair color")
		t.Fail()
	}
	// Missing #
	if validHairColor("0099af") {
		t.Log("Error, 0099af considered a valid hair color")
		t.Fail()
	}
}

func TestValidEyeColor(t *testing.T) {
	if !validEyeColor("grn") {
		t.Log("Error, grn not considered a valid eye color")
		t.Fail()
	}
	if validEyeColor("green") {
		t.Log("Error, green considered a valid eye color")
		t.Fail()
	}
	if validEyeColor("grns") {
		t.Log("Error, grns considered a valid eye color")
		t.Fail()
	}
}

func TestValidPassportNumber(t *testing.T) {
	if !validPassportNumber("000000001") {
		t.Log("Error, 000000001 not considered a valid passport number")
		t.Fail()
	}
	// Too short
	if validPassportNumber("12345678") {
		t.Log("Error, 12345678 considered a valid passport number")
		t.Fail()
	}
	// Too long
	if validPassportNumber("1234567890") {
		t.Log("Error, 1234567890 considered a valid passport number")
		t.Fail()
	}
	// Not a number
	if validPassportNumber("a1234567890") {
		t.Log("Error, a1234567890 considered a valid passport number")
		t.Fail()
	}
}
