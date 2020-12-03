package main

import (
	"bufio"
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"os"
	"path/filepath"
)

const TREE uint8 = '#'

func parseInput(filename string) []string {
	lines := make([]string, 0)

	absFilePath, absFilePathErr := filepath.Abs(filename)
	lib.Check(absFilePathErr)

	file, err := os.Open(absFilePath)
	lib.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	lib.Check(scanner.Err())
	return lines
}

func partOne() int {
	count := 0
	parsedInput := parseInput("day3/input.txt")

	for index, line := range parsedInput {
		pos := (index * 3) % len(line)
		isTree := line[pos] == TREE
		if isTree {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(partOne())
}
