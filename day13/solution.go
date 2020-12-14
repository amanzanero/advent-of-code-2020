package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/day13/utils"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"math"
	"strconv"
)

func main() {
	parsedLines := lib.ParseLines("day13/input.txt")
	earliestDeparture, err := strconv.Atoi(parsedLines[0])
	lib.Check(err)
	ids := utils.ParseIds(parsedLines[1])

	partOne(earliestDeparture, ids)
	partTwo(ids)
}

func partOne(earliestDeparture int, ids []string) {
	defer lib.Elapsed("-- took: ")()
	minDiff := math.MaxInt64
	smallestId := 0
	for _, id := range ids {
		if id == "x" {
			continue
		}
		curr, err := strconv.Atoi(id)
		lib.Check(err)
		nextDepart := (earliestDeparture/curr)*curr + curr
		diff := nextDepart - earliestDeparture
		if diff < minDiff {
			minDiff = diff
			smallestId = curr
		}
	}
	fmt.Printf("Smallest id: %d\n", smallestId*minDiff)
}

func partTwo(ids []string) {
	defer lib.Elapsed("-- took: ")()

	modulos := make([]int, 0)
	remainders := make([]int, 0)
	product := 1
	for i, inputId := range ids {
		if inputId == "x" {
			continue
		} else {
			id, err := strconv.Atoi(inputId)
			lib.Check(err)
			product *= id
			modulos = append(modulos, id)
			remainders = append(remainders, id-i)
		}
	}

	result := 0
	for i := 0; i < len(modulos); i++ {
		pExceptI := product / modulos[i]
		result += remainders[i] * utils.ModuloInverse(pExceptI, modulos[i]) * pExceptI
	}
	fmt.Printf("Earliest consecutive: %d\n", result%product)
}
