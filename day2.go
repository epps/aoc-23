package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	RED_CUBES   int64 = 12
	GREEN_CUBES int64 = 13
	BLUE_CUBES  int64 = 14
)

var cubeMap = map[string]int64{
	"red":   RED_CUBES,
	"green": GREEN_CUBES,
	"blue":  BLUE_CUBES,
}

type Day2Puzzle struct {
	input  *os.File
	number int
}

func (d *Day2Puzzle) Solve() (interface{}, error) {
	if d.number == PUZZLE_TWO {
		return nil, fmt.Errorf("puzzle 2 of day 2 not implemented")
	}
	scanner := bufio.NewScanner(d.input)
	scanner.Split(bufio.ScanLines)

	gameIds := make([]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()

		lineBrokenOnColon := strings.Split(line, ":")
		gameStr := strings.Split(lineBrokenOnColon[0], " ")[1]
		rounds := strings.Split(lineBrokenOnColon[1], ";")

		isValidGame := true
		for _, round := range rounds {
			cubes := strings.Split(round, ",")
			for _, cube := range cubes {
				countParts := strings.Split(strings.TrimSpace(cube), " ")
				count, err := strconv.ParseInt(countParts[0], 0, 0)
				if err != nil {
					log.Fatalf("Failed to parse cube count %s due to error: %v", countParts[0], err)
				}
				maxCountForColor := cubeMap[countParts[1]]

				if count > maxCountForColor {
					isValidGame = false
					break
				}
			}
		}

		if isValidGame {
			gameId, err := strconv.ParseInt(gameStr, 0, 0)
			if err != nil {
				log.Fatalf("Failed to parse game %s due to error: %v", gameStr, err)
			}
			gameIds = append(gameIds, gameId)
		}
	}

	var sum int64

	for _, id := range gameIds {
		sum += id
	}

	return sum, nil
}

func NewDay2Puzzle(puzzleNumber int) (Puzzle, error) {
	f, err := os.Open("./puzzle-inputs/day2.txt")
	if err != nil {
		return nil, err
	}

	return &Day2Puzzle{
		input:  f,
		number: puzzleNumber,
	}, nil
}
