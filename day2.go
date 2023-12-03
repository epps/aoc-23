package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(d.input)
	scanner.Split(bufio.ScanLines)

	results := make([]int64, 0)
	for scanner.Scan() {
		line := scanner.Text()

		switch d.number {
		case PUZZLE_ONE:
			d.processLineForPuzzle1(line, &results)
		case PUZZLE_TWO:
			d.processLineForPuzzle2(line, &results)
		default:
			log.Fatalf("unrecognized puzzle number: %d", d.number)
		}
	}

	var sum int64

	for _, id := range results {
		sum += id
	}

	return sum, nil
}

func (d *Day2Puzzle) processLineForPuzzle1(line string, results *[]int64) {
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
		*results = append(*results, gameId)
	}
}

func (d *Day2Puzzle) processLineForPuzzle2(line string, results *[]int64) {
	lineBrokenOnColon := strings.Split(line, ":")
	// gameStr := strings.Split(lineBrokenOnColon[0], " ")[1]
	rounds := strings.Split(lineBrokenOnColon[1], ";")

	var maxRed int64
	var maxGreen int64
	var maxBlue int64
	for _, round := range rounds {
		cubes := strings.Split(round, ",")
		for _, cube := range cubes {
			countParts := strings.Split(strings.TrimSpace(cube), " ")
			count, err := strconv.ParseInt(countParts[0], 0, 0)
			if err != nil {
				log.Fatalf("Failed to parse cube count %s due to error: %v", countParts[0], err)
			}
			color := countParts[1]
			switch color {
			case "red":
				if count > maxRed {
					maxRed = count
				}
			case "green":
				if count > maxGreen {
					maxGreen = count
				}
			case "blue":
				if count > maxBlue {
					maxBlue = count
				}
			default:
				log.Fatalf("unrecognized color %s", color)
			}
		}
	}

	power := maxRed * maxGreen * maxBlue

	*results = append(*results, power)
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
