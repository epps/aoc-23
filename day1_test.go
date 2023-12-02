package main

import "testing"

func TestProcessLineForPuzzle2(t *testing.T) {
	cases := []struct {
		line           string
		expectedDigits []string
	}{
		{
			line:           "eightwothree",
			expectedDigits: []string{"8", "3"},
		},
		{
			line:           "trsdgcxcseven39dpmzs",
			expectedDigits: []string{"7", "9"},
		},
		{
			line:           "hnjcrxeightonejnlvm4hstmcsevensix",
			expectedDigits: []string{"8", "6"},
		},
		{
			line:           "7xv3one",
			expectedDigits: []string{"7", "1"},
		},
		{
			line:           "eighttkbtzjz6nineeight",
			expectedDigits: []string{"8", "8"},
		},
		{
			line:           "5knjbxgvhktvfcq89onefive",
			expectedDigits: []string{"5", "5"},
		},
		{
			line:           "two9eightxnpdj61kzcdpnpnpfgsdrbcflh",
			expectedDigits: []string{"2", "1"},
		},
		{
			line:           "sevenrdmhnldsmdnineqfrgjhmhnnqkztxzm7",
			expectedDigits: []string{"7", "7"},
		},

		{
			line:           "two1nine",
			expectedDigits: []string{"2", "9"},
		},
		{
			line:           "abcone2threexyz",
			expectedDigits: []string{"1", "3"},
		},
		{
			line:           "xtwone3four",
			expectedDigits: []string{"2", "4"},
		},
		{
			line:           "4nineeightseven2",
			expectedDigits: []string{"4", "2"},
		},
		{
			line:           "zoneight234",
			expectedDigits: []string{"1", "4"},
		},
		{
			line:           "7pqrstsixteen",
			expectedDigits: []string{"7", "6"},
		},
	}

	puzzle := Day1Puzzle{}

	for _, c := range cases {
		var first string
		var last string
		puzzle.processLineForPuzzle2(c.line, &first, &last)

		if first != c.expectedDigits[0] || last != c.expectedDigits[1] {
			t.Fatalf("Expected line %s to yield digits %v but received %s and %s", c.line, c.expectedDigits, first, last)
		}
	}
}
