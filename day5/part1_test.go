package main

import (
	"testing"
)

type TestCase struct {
	input    string
	expected int
}

func TestRow(t *testing.T) {

	testCases := []TestCase{
		{
			input:    "FFFFFFF",
			expected: 0,
		},
		{
			input:    "FBFBBFF",
			expected: 44,
		},
		{
			input:    "BBBBBBB",
			expected: 127,
		},
	}

	for _, test := range testCases {

		actual := GetRow(test.input)

		if actual != test.expected {
			t.Errorf("Got %d, expected %d", actual, test.expected)
		}
	}
}

func TestColumn(t *testing.T) {

	testCases := []TestCase{
		{
			input:    "RRR",
			expected: 7,
		},
		{
			input:    "RLR",
			expected: 5,
		},
		{
			input:    "LLL",
			expected: 0,
		},
	}

	for _, test := range testCases {
		actual := GetColumn(test.input)

		if actual != test.expected {
			t.Errorf("Got %d, expected %d", actual, test.expected)
		}
	}
}
