package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/day14/memory"
	"github.com/amanzanero/advent-of-code-2020/day14/utils"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"strings"
)

func main() {
	parsedLines := lib.ParseLines("day14/input.txt")
	instructions := parseMemoryInstructions(parsedLines)
	partOne(instructions)
	partTwo(instructions)
}

func partOne(instructions []*memory.MemExpression) {
	defer lib.Elapsed("-- took: ")()
	memblock := make(memory.MemBlock)
	for _, instruction := range instructions {
		memblock[instruction.Destination] = (instruction.Value & (^instruction.Mask)) | instruction.MaskValue
	}

	var sum uint64 = 0
	for _, value := range memblock {
		sum += value
	}
	fmt.Printf("Sum of memblock values: %d\n", sum)
}

func partTwo(instructions []*memory.MemExpression) {
	defer lib.Elapsed("-- took: ")()
	memblock := make(memory.MemBlock)

	var address uint64
	cache := make(map[uint64][]uint64)
	var bitCombos []uint64
	for _, instruction := range instructions {
		address = instruction.Destination | instruction.MaskValue
		value, exists := cache[instruction.MaskValue]
		if exists {
			bitCombos = value
		} else {
			bitCombos = instruction.WildcardCombinations()
		}

		for _, combo := range bitCombos {
			address = combo | (address & ^(instruction.Wildcard))
			memblock[address] = instruction.Value
		}
	}

	var sum uint64 = 0
	for _, value := range memblock {
		sum += value
	}
	fmt.Printf("Sum of memblock values: %d\n", sum)
}

func parseMemoryInstructions(lines []string) []*memory.MemExpression {
	instructions := make([]*memory.MemExpression, 0)
	for iter := 0; iter < len(lines); {
		mask, value, wildcard := utils.ParseMask(lines[iter])
		iter++
		for iter < len(lines) && strings.HasPrefix(lines[iter], "mem") {
			exp := utils.ParseExpression(lines[iter])

			exp.Mask = mask
			exp.MaskValue = value
			exp.Wildcard = wildcard
			instructions = append(instructions, exp)
			iter++
		}
	}
	return instructions
}
