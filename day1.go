package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

type Day1Puzzle1 struct {
	input *os.File
}

func (d *Day1Puzzle1) Solve() (interface{}, error) {
	log.Println("Solving Day 1 Puzzle 1 ...")
	defer d.input.Close()

	scanner := bufio.NewScanner(d.input)
	scanner.Split(bufio.ScanLines)

	coordinates := make([]int64, 0)
	for scanner.Scan() {
		var firstDigit string
		var lastDigit string
		text := scanner.Text()

		for _, char := range text {
			if unicode.IsNumber(rune(char)) {
				if firstDigit == "" {
					firstDigit = string(char)
				}
				lastDigit = string(char)
			}
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

func NewDay1Puzzle1() (Puzzle, error) {
	f, err := os.Open("./puzzle-inputs/day1.txt")
	if err != nil {
		return nil, err
	}

	return &Day1Puzzle1{
		input: f,
	}, nil
}
