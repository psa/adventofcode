package main

import (
	"reflect"
	"testing"
)

func TestParsePasswordLines(t *testing.T) {
	var passwordLines = []string{
		"1-10 j: vrfjljjwbsv",
		"1-2 broken"}
	var expected = []passwordData{
		passwordData{
			Min:       1,
			Max:       10,
			Character: "j",
			Password:  "vrfjljjwbsv",
		},
	}
	result := parsePasswordLines(passwordLines)

	if !reflect.DeepEqual(expected, result) {
		t.Log("Error, password lines wrong", result)
		t.Fail()
	}
}

func TestScanPasswords(t *testing.T) {
	var data = []passwordData{
		passwordData{ // Min (correct)
			Min:       1,
			Max:       10,
			Character: "x",
			Password:  "x",
		},
		passwordData{ // Max (correct)
			Min:       2,
			Max:       2,
			Character: "x",
			Password:  "xx",
		},
		passwordData{ // More than max
			Min:       1,
			Max:       2,
			Character: "x",
			Password:  "xxxxx",
		},
		passwordData{ // Less than Min
			Min:       2,
			Max:       2,
			Character: "x",
			Password:  "x",
		},
	}
	result := scanPasswords(data)

	if result != 2 {
		t.Log("Error, expect 2 correct results, got", result)
		t.Fail()
	}
}

func TestScanPasswordsNewPolicy(t *testing.T) {
	var data = []passwordData{
		passwordData{ // Valid, position 1 contains a and position 3 does not
			Min:       1,
			Max:       3,
			Character: "a",
			Password:  "abcde",
		},
		passwordData{ // Invalid, neither position 1 nor position 3 contains b
			Min:       1,
			Max:       3,
			Character: "b",
			Password:  "cdefg",
		},
		passwordData{ // Invalid, both position 2 and position 9 contain c.
			Min:       2,
			Max:       9,
			Character: "c",
			Password:  "ccccccccc",
		},
		passwordData{ // Invalid, neither position is in the password
			Min:       2,
			Max:       9,
			Character: "c",
			Password:  "c",
		},
		passwordData{ // Invalid, Min incorrect and Max outside
			Min:       2,
			Max:       9,
			Character: "c",
			Password:  "ddd",
		},
		passwordData{ // Valid, Min correct and Max outside
			Min:       2,
			Max:       9,
			Character: "c",
			Password:  "ccc",
		},
	}
	result := scanPasswordsNewPolicy(data)

	if result != 2 {
		t.Log("Error, expect 1 correct result, got", result)
		t.Fail()
	}
}
