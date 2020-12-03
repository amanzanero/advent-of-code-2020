package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/lib"
)

const MAGICNUMBER = 2020

func partOne(intInput []int) int {

	seenNums := make(map[int]bool) // New empty set

	for _, num := range intInput {
		// found a solution
		if compliment := MAGICNUMBER - num; seenNums[compliment] {
			return num * compliment
		}
		// insert into seen
		seenNums[num] = true
	}
	return -1
}

func partTwo(intInput []int) int {
	// do one linear pass to get sum of two compliments
	firstComplimentPass := make(map[int]int)
	for _, num := range intInput {
		compliment := MAGICNUMBER - num
		firstComplimentPass[compliment] = num
	}

	// now we basically do part one for every element
	for compliment, num := range firstComplimentPass {

		seenNums := make(map[int]bool) // New empty set
		for _, secondNum := range intInput {
			// found a solution
			if secondCompliment := compliment - secondNum; seenNums[secondCompliment] {
				return num * secondNum * secondCompliment
			}
			// insert into seen
			seenNums[secondNum] = true
		}

	}

	return -1
}

func main() {
	intInput := lib.GetIntArrayInput("day1/input.txt")

	fmt.Print("D1P1: ")
	fmt.Println(partOne(intInput))

	fmt.Print("D1P1: ")
	fmt.Println(partTwo(intInput))
}
