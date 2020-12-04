package main

import (
	"bufio"
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"os"
	"path/filepath"
	"sync"
)

const TREE uint8 = '#'

type Slope struct {
	Right int
	Down  int
}

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
	parsedInput := parseInput("day3/input.txt")

	slope := Slope{3, 1}

	return countTrees(&slope, parsedInput)
}

func partTwo() int {
	parsedInput := parseInput("day3/input.txt")
	counts := [5]int{0, 0, 0, 0, 0}
	slopes := [5]Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	wg := sync.WaitGroup{}
	for index, slope := range slopes {
		wg.Add(1)
		go func(ind int, s Slope) {
			counts[ind] = countTrees(&s, parsedInput)
			wg.Done()
		}(index, slope)
	}
	wg.Wait()
	ans := 1
	for _, num := range counts {
		ans *= num
	}
	return ans
}

func countTrees(slope *Slope, lines []string) int {
	count := 0
	for row, iter := 0, 0; row < len(lines); row += slope.Down {
		line := lines[row]
		column := (slope.Right * iter) % len(line)
		if line[column] == TREE {
			count++
		}
		iter++
	}
	return count
}

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}
