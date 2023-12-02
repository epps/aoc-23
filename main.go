package main

import (
	"flag"
	"fmt"
	"log"
)

type PuzzleNumber int

const (
	PUZZLE_ONE = 1
	PUZZLE_TWO = 2
)

type Puzzle interface {
	Solve() (interface{}, error)
}

func main() {
	fmt.Println("Advent of Code 2023!")

	dayNum := flag.Int("day", 1, "selects a day")
	puzzleNum := flag.Int("puzzle", 1, "selects a day's puzzle")

	flag.Parse()

	if *dayNum <= 0 || *dayNum > 25 {
		log.Fatalf("day must be an integer between 1 and 25; received %d", *dayNum)
	}

	if *puzzleNum <= 0 || *puzzleNum >= 3 {
		log.Fatalf("puzzle must be either a 1 or a 2; received %d", *puzzleNum)
	}

	var puzzle Puzzle
	var puzzleErr error
	switch *dayNum {
	case 1:
		puzzle, puzzleErr = NewDay1Puzzle(*puzzleNum)
	case 2:
		puzzle, puzzleErr = NewDay2Puzzle(*puzzleNum)
	default:
		puzzleErr = fmt.Errorf("day %d not implemented", *dayNum)
	}
	if puzzleErr != nil {
		log.Fatalf("failed to get puzzle %d for day %d: %v", *puzzleNum, *dayNum, puzzleErr)
	}

	solution, err := puzzle.Solve()

	if err != nil {
		log.Fatalf("failed to solve puzzle %d for day %d: %v", *puzzleNum, *dayNum, err)
	}

	log.Printf("Solution for puzzle %d of day %d is %v (type %T)", *puzzleNum, *dayNum, solution, solution)
}
