package main

import (
	"bufio"
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type Line struct {
	min      int
	max      int
	letter   string
	password string
}

func parseLine(line string) Line {
	delimiter := regexp.MustCompile(`[-: ]`)
	splitLine := delimiter.Split(line, -1)

	min, minErr := strconv.Atoi(splitLine[0])
	max, maxErr := strconv.Atoi(splitLine[1])
	lib.Check(minErr)
	lib.Check(maxErr)

	letter := splitLine[2]
	password := splitLine[4]
	return Line{min: min, max: max, letter: letter, password: password}
}

func parseInput(fileName string) []Line {
	lines := make([]Line, 0)

	absFilePath, absFilePathErr := filepath.Abs(fileName)
	lib.Check(absFilePathErr)

	file, err := os.Open(absFilePath)
	defer file.Close()
	lib.Check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := parseLine(scanner.Text())
		lines = append(lines, line)
	}
	lib.Check(scanner.Err())
	return lines
}

func partOne() int {
	parsedInput := parseInput("day2/input.txt")
	correctCount := 0

	for _, line := range parsedInput {
		count := strings.Count(line.password, line.letter)
		if count <= line.max && count >= line.min {
			correctCount += 1
		}
	}

	return correctCount
}

func existsAt(source, letter string, pos int) bool {
	letterChar := letter[0]
	return pos <= len(source)-1 && source[pos] == letterChar
}

func partTwo() int {
	parsedInput := parseInput("day2/input.txt")
	correctCount := 0

	for _, line := range parsedInput {
		firstPosExists := existsAt(line.password, line.letter, line.min-1)
		secondPosExists := existsAt(line.password, line.letter, line.max-1)

		if (firstPosExists || secondPosExists) && firstPosExists != secondPosExists {
			correctCount++
		}
	}

	return correctCount
}

func main() {
	fmt.Print("D2P1: ")
	fmt.Println(partOne())

	fmt.Print("D2P2: ")
	fmt.Println(partTwo())
}
