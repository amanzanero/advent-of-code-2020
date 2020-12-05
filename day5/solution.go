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

func partTwo() int {
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
			return int(currId - 1)
		}
	}

	// failed
	return -1
}

func parseSeatId(sequence string) float64 {
	row := binaryBoarding(sequence[:7], 'F')
	col := binaryBoarding(sequence[7:], 'L')
	return row*8 + col
}

func binaryBoarding(sequence string, comp uint8) float64 {
	bits := math.Pow(2, float64(len(sequence)))

	start, end := float64(0), bits-1

	var midpoint float64

	for i := 0; i < len(sequence); i++ {
		char := sequence[i]
		midpoint = (start + end) / 2
		if char == comp {
			end = math.Floor(midpoint)
		} else {
			start = math.Ceil(midpoint)
		}
	}

	// can be start or end by this point since start == end
	return start
}
