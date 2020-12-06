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

func partOne() float64 {
	parsedLines := lib.ParseLines("day5/input.txt")
	var max float64 = -1

	for _, line := range parsedLines {
		seatId := parseSeatId(line)
		max = math.Max(seatId, max)
	}
	return max
}

func partTwo() float64 {
	parsedLines := lib.ParseLines("day5/input.txt")

	// parse all ids
	seatIds := make([]float64, 0)
	for _, line := range parsedLines {
		seatId := parseSeatId(line)
		seatIds = append(seatIds, seatId)
	}
	sort.Float64s(seatIds)

	// find missing id
	for index, currId := range seatIds[1:] {
		prevId := seatIds[index]
		isNotContiguous := currId-prevId > 1

		if isNotContiguous {
			return currId - 1
		}
	}

	// failed
	return -1
}

func parseSeatId(sequence string) float64 {
	row := evaluateBinary(sequence[:7], 'B')
	col := evaluateBinary(sequence[7:], 'R')
	return float64(row*8 + col)
}

func evaluateBinary(sequence string, comp uint8) int {
	ans, lastIndex := 0, len(sequence)-1

	for i := lastIndex; i >= 0; i-- {
		if sequence[i] == comp {
			ans |= 1 << (lastIndex - i)
		}
	}
	return ans
}
