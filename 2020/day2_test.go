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
