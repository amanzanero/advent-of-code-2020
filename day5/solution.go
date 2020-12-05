package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"math"
	"sort"
)

func main() {
	fmt.Println(partOne())
	fmt.Println(partTwo())
}

func partOne() int {
	parsedLines := lib.ParseLines("day5/input.txt")
	max := math.MinInt64

	for _, line := range parsedLines {
		row := binaryBoarding(line[:7], 'F')
		col := binaryBoarding(line[7:], 'L')

		seatId := row*8 + col
		max = int(math.Max(float64(seatId), float64(max)))
	}
	return max
}

func partTwo() int {
	parsedLines := lib.ParseLines("day5/input.txt")

	// parse all ids
	sids := make([]int, 0)
	for _, line := range parsedLines {
		row := binaryBoarding(line[:7], 'F')
		col := binaryBoarding(line[7:], 'L')

		seatId := row*8 + col
		sids = append(sids, seatId)
	}
	sort.Ints(sids)

	// find missing id
	for index, currId := range sids[1:] {
		prevId := sids[index]
		if currId-prevId > 1 {
			return currId - 1
		}
	}

	// failed
	return -1
}

func binaryBoarding(sequence string, comp uint8) int {
	bits := math.Pow(2, float64(len(sequence)))

	start, end := 0, int(bits)-1

	var midpoint int

	for i := 0; i < len(sequence); i++ {
		char := sequence[i]
		midpoint = (start + end) / 2
		if char == comp {
			end = midpoint
		} else {
			start = midpoint + 1
			midpoint = start
		}
	}
	return midpoint
}
