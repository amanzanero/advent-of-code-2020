package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/day11/ferry"
	"github.com/amanzanero/advent-of-code-2020/lib"
)

func main() {
	rows := lib.ParseLines("day11/input.txt")
	partOne(rows)
}

func partOne(rows []string) {
	stop1 := lib.Elapsed("-- took: ")
	count := simulateSeating(ferry.PersonCanSit, ferry.PersonWillLeave, rows)
	fmt.Printf("Day 1 part 1: %d\n", count)
	stop1()

	stop2 := lib.Elapsed("-- took: ")
	count = simulateSeating(ferry.PersonCanSitFirstEachDirection, ferry.TolerantPersonWillLeave, rows)
	fmt.Printf("Day 1 part 2: %d\n", count)
	stop2()
}

func simulateSeating(emptySeatPolicy, takenSeatPolicy func(int, int, [][]uint8) bool, rows []string) int {
	occupiedSeats := 0
	seatingChartChanged := true
	cache1 := ferry.NbyMMatrix(len(rows), len(rows[0]))
	cache2 := ferry.NbyMMatrix(len(rows), len(rows[0]))

	// copy rows to cache1
	for i := 0; i < len(rows); i++ {
		for j := 0; j < len(rows[i]); j++ {
			cache1[i][j] = rows[i][j]
		}
	}

	for iter := 0; seatingChartChanged || iter == 0; iter++ {
		seatingChartChanged = false
		occupiedSeats = 0

		var currentSeats [][]uint8
		var nextSeats [][]uint8
		if iter%2 == 0 {
			currentSeats = cache1
			nextSeats = cache2
		} else {
			currentSeats = cache2
			nextSeats = cache1
		}

		for x := 0; x < len(currentSeats); x++ {
			for y := 0; y < len(currentSeats[x]); y++ {
				state := currentSeats[x][y]
				switch state {
				case ferry.EMPTY:
					if emptySeatPolicy(x, y, currentSeats) {
						nextSeats[x][y] = ferry.SEAT
						occupiedSeats++
						seatingChartChanged = true
					} else {
						nextSeats[x][y] = ferry.EMPTY
					}
					break
				case ferry.SEAT:
					if takenSeatPolicy(x, y, currentSeats) {
						nextSeats[x][y] = ferry.EMPTY
						seatingChartChanged = true
					} else {
						nextSeats[x][y] = ferry.SEAT
						occupiedSeats++
					}
					break
				case ferry.FLOOR:
					nextSeats[x][y] = ferry.FLOOR
					break
				}
			}
		}
	}
	return occupiedSeats
}
