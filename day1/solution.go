package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/lib"
)

const magicNumber = 2020

func partOne(intInput []int) int {

	seenNums := make(map[int]bool) // New empty set

	for _, num := range intInput {
		// found a solution
		if compliment := magicNumber - num; seenNums[compliment] {
			return num * compliment
		}
		// insert into seen
		seenNums[num] = true
	}
	return -1
}

func main() {
	intInput := lib.GetIntArrayInput("day1/input.txt")

	fmt.Print("D1P1: ")
	fmt.Println(partOne(intInput))

}
