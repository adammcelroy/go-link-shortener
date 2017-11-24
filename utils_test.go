package main

import (
	"errors"
	"testing"
)

func printDiscrepancy(t *testing.T, expected string, got string) {
	t.Errorf("\nExpected: \n%v \nGot: \n%v", expected, got)
}

func TestEncodeURL(t *testing.T) {
	input := "https://google.com/search?q=golang"
	expected := "https%3A%2F%2Fgoogle.com%2Fsearch%3Fq%3Dgolang"
	output := encodeURL(input)

	if output != expected {
		t.Errorf("\n", "Expected URL to be encoded properly")
		printDiscrepancy(t, expected, output)
	}
}

func TestEnforceProtocol(t *testing.T) {
	input1 := "google.com"
	input2 := "http://facebook.com"
	input3 := "https://twitter.com"
	expected1 := "http://google.com"
	expected2 := "http://facebook.com"
	expected3 := "https://twitter.com"
	output1 := enforceProtocol(input1)
	output2 := enforceProtocol(input2)
	output3 := enforceProtocol(input3)

	if output1 != expected1 {
		t.Error("\n", "Expected URL without protocol to be prefixed with 'http://'")
		printDiscrepancy(t, expected1, output1)
	}
	if output2 != expected2 {
		t.Error("\n", "Expected URL with existing 'http://' protocol to be left alone")
		printDiscrepancy(t, expected2, output2)
	}
	if output3 != expected3 {
		t.Error("\n", "Expected URL with existing 'https://' protocol to be left alone")
		printDiscrepancy(t, expected3, output3)
	}
}

func TestPanicIfError(t *testing.T) {
	err := errors.New("Some error")

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic if error detected")
		}
	}()

	panicIfError(err)
}
