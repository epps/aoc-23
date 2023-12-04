package main

import (
	"os"
	"testing"
)

func TestDay3Puzzle1(t *testing.T) {
	testFile, err := os.Open("./puzzle-inputs/day3_test.txt")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}

	puzzle := Day3Puzzle{
		input:  testFile,
		number: 1,
	}

	solution, err := puzzle.Solve()
	if err != nil {
		t.Fatalf("Failed to solve puzzle: %v", err)
	}

	if solution.(int64) != 4361 {
		t.Fatalf("Expected solution to be 4361 but received %d", solution.(int64))
	}
}
