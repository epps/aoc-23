package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Day3Puzzle struct {
	input  *os.File
	number int
}

func (d *Day3Puzzle) Solve() (interface{}, error) {
	scanner := bufio.NewScanner(d.input)
	scanner.Split(bufio.ScanLines)

	results := make([]int64, 0)
	// need to populate previous line with an empty slice
	prevRow := make([]rune, 0)
	// Read the first line to populate the current row
	if hasLine := scanner.Scan(); !hasLine {
		log.Fatal("Day 3 input file is empty")
	}
	currentRow := []rune(scanner.Text())

	for scanner.Scan() {
		nextRow := []rune(scanner.Text())

		switch d.number {
		case PUZZLE_ONE:
			d.processLinesForPuzzle1(prevRow, currentRow, nextRow, &results)
		case PUZZLE_TWO:
			log.Fatal("puzzle 2 for day 3 not implemented")
		default:
			log.Fatalf("unrecognized puzzle number: %d", d.number)
		}

		prevRow = currentRow
		currentRow = nextRow
	}

	// process the last line
	switch d.number {
	case PUZZLE_ONE:
		d.processLinesForPuzzle1(prevRow, currentRow, []rune{}, &results)
	case PUZZLE_TWO:
		log.Fatal("puzzle 2 for day 3 not implemented")
	default:
		log.Fatalf("unrecognized puzzle number: %d", d.number)
	}

	var sum int64

	for _, val := range results {
		sum += val
	}

	return sum, nil
}

func (d *Day3Puzzle) isSymbol(char rune) bool {
	return char != rune('.') && !unicode.IsNumber(char)
}

func (d *Day3Puzzle) processLinesForPuzzle1(prevRow, currentRow, nextRow []rune, results *[]int64) {

	for i := 0; i < len(currentRow); i++ {
		if currentRow[i] == rune('.') {
			continue
		}
		// If we encounter a number, there several ways for it to be adjacent to a symbol
		// * left of start index
		// * right of end index
		// * up and left of start
		// * down and left of start
		// * up/down from any index "inside" of the number
		// * up and right of end index
		// * down and right of end index
		if unicode.IsNumber(currentRow[i]) {
			start := i
			number := string(currentRow[i])
			for {
				if i+1 < len(currentRow) && unicode.IsNumber(currentRow[i+1]) {
					number += string(currentRow[i+1])
					i += 1
				} else {
					break
				}
			}
			hasAdjacentSymbol := false

			if start-1 >= 0 && d.isSymbol(currentRow[start-1]) { // left of start
				hasAdjacentSymbol = true
			} else if i+1 < len(currentRow) && d.isSymbol(currentRow[i+1]) { // right of end
				hasAdjacentSymbol = true
			} else if len(prevRow) > 0 && start-1 >= 0 && d.isSymbol(prevRow[start-1]) { // up and left of start
				hasAdjacentSymbol = true
			} else if start-1 >= 0 && len(nextRow) > 0 && d.isSymbol(nextRow[start-1]) { // down and left of start
				hasAdjacentSymbol = true
			} else if len(prevRow) > 0 && i+1 < len(prevRow) && d.isSymbol(prevRow[i+1]) { // up and right of end
				hasAdjacentSymbol = true
			} else if i+1 < len(nextRow) && len(nextRow) > 0 && d.isSymbol((nextRow[i+1])) { // down and right of end
				hasAdjacentSymbol = true
			} else {
				for {
					// iterate over the number digits
					if start == i+1 {
						break
					}

					// check directly above and below the digits of the number
					if len(nextRow) > 0 && d.isSymbol(nextRow[start]) || (len(prevRow) > 0 && d.isSymbol(prevRow[start])) {
						hasAdjacentSymbol = true
						break
					}
					start += 1
				}
			}

			if hasAdjacentSymbol {
				partNumber, err := strconv.ParseInt(number, 0, 0)
				if err != nil {
					log.Fatalf("Failed to parse part number %s: %v", number, err)
				}
				*results = append(*results, partNumber)
			}
		}
	}
}

func NewDay3Puzzle(puzzleNumber int) (Puzzle, error) {
	f, err := os.Open("./puzzle-inputs/day3.txt")
	if err != nil {
		return nil, err
	}

	return &Day3Puzzle{
		input:  f,
		number: puzzleNumber,
	}, nil
}
