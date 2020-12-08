package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/day8/vmachine"
	"github.com/amanzanero/advent-of-code-2020/lib"
)

func main() {
	parsedOps := lib.ParseLines("day8/input.txt")

	vm := vmachine.CreateVMachine()
	for _, parsed := range parsedOps {
		op := vmachine.ParseOpFromString(parsed)
		vm.AddOp(op)
	}

	fmt.Printf("Day 8 part 1: %d\n", partOne(vm))
	vm.Reset()
	fmt.Printf("Day 8 part 2: %d\n", partTwo(vm))

}

func partOne(vm *vmachine.VMachine) int {
	// run the program
	vm.RunProgram()
	return vm.Acc()
}

func partTwo(vm *vmachine.VMachine) int {
	for index, op := range vm.Ops {
		var replacedCopy *vmachine.VMachine
		var replace func(int) = nil

		if op.OpType() == "nop" {
			replacedCopy = vm.Copy()
			replace = replacedCopy.ReplaceJmpAt
		} else if op.OpType() == "jmp" {
			replacedCopy = vm.Copy()
			replace = replacedCopy.ReplaceNOpAt
		}

		// swap and make a copy
		if replace != nil && replacedCopy != nil {
			replace(index)
			replacedCopy.RunProgram()
			if !replacedCopy.IsCorrupted() {
				return replacedCopy.Acc()
			}
		}
	}

	return -1
}
