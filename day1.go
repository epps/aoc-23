package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Day1Puzzle struct {
	input  *os.File
	number int
}

func (d *Day1Puzzle) Solve() (interface{}, error) {
	log.Printf("Solving Day 1 Puzzle %d ...", d.number)
	defer d.input.Close()

	scanner := bufio.NewScanner(d.input)
	scanner.Split(bufio.ScanLines)

	coordinates := make([]int64, 0)
	for scanner.Scan() {
		var firstDigit string
		var lastDigit string
		text := scanner.Text()

		switch d.number {
		case PUZZLE_ONE:
			d.processLineForPuzzle1(text, &firstDigit, &lastDigit)
		case PUZZLE_TWO:
			d.processLineForPuzzle2(text, &firstDigit, &lastDigit)
		default:
			log.Panicf("unrecognized puzzle number: %d", d.number)
		}

		coordinate, err := strconv.ParseInt(firstDigit+lastDigit, 0, 0)
		if err != nil {
			log.Panicf("failed to parse %s: %v", firstDigit+lastDigit, err)
		}
		coordinates = append(coordinates, coordinate)
	}

	var sum int64

	for _, coord := range coordinates {
		sum += coord
	}

	return sum, nil
}

func (d *Day1Puzzle) processLineForPuzzle1(line string, firstDigit *string, lastDigit *string) {
	for _, char := range line {
		if unicode.IsNumber(rune(char)) {
			if *firstDigit == "" {
				*firstDigit = string(char)
			}
			*lastDigit = string(char)
		}
	}
}

func (d *Day1Puzzle) processLineForPuzzle2(line string, firstDigit *string, lastDigit *string) {
	chars := []rune(line)
	idx := 0
	for {
		if idx >= len(chars) {
			break
		}
		var digit string
		char := chars[idx]
		if unicode.IsNumber(rune(char)) {
			digit = string(char)
		} else {
			var digitChars []rune
			var parsedDigit string

			switch char {
			case rune('o'):
				digitChars = []rune("one")
				parsedDigit = "1"
			case rune('t'):
				if idx+1 >= len(chars) {
					break
				}
				nextChar := chars[idx+1]
				if nextChar == rune('w') {
					digitChars = []rune("two")
					parsedDigit = "2"
				} else if nextChar == rune('h') {
					digitChars = []rune("three")
					parsedDigit = "3"
				}
			case rune('f'):
				if idx+1 >= len(chars) {
					break
				}
				nextChar := chars[idx+1]
				if nextChar == rune('o') {
					digitChars = []rune("four")
					parsedDigit = "4"
				} else if nextChar == rune('i') {
					digitChars = []rune("five")
					parsedDigit = "5"
				}
			case rune('s'):
				if idx+1 >= len(chars) {
					break
				}
				nextChar := chars[idx+1]
				if nextChar == rune('i') {
					digitChars = []rune("six")
					parsedDigit = "6"
				} else if nextChar == rune('e') {
					digitChars = []rune("seven")
					parsedDigit = "7"
				}
			case rune('e'):
				digitChars = []rune("eight")
				parsedDigit = "8"
			case rune('n'):
				digitChars = []rune("nine")
				parsedDigit = "9"
			}
			if parsedDigit != "" {
				isDigit := true
				for i, c := range digitChars {
					if idx+1 >= len(chars) || c != chars[idx+i] {
						isDigit = false
						break
					}
				}
				if isDigit {
					idx += len(digitChars) - 1
					digit = parsedDigit
				}
			}
		}
		if *firstDigit == "" && digit != "" {
			*firstDigit = digit
		}
		if digit != "" {
			*lastDigit = digit
		}
		idx += 1
	}
}

func NewDay1Puzzle(puzzleNumber int) (Puzzle, error) {
	f, err := os.Open("./puzzle-inputs/day1.txt")
	if err != nil {
		return nil, err
	}

	return &Day1Puzzle{
		input:  f,
		number: puzzleNumber,
	}, nil
}
