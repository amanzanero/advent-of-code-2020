package vmachine

import (
	"github.com/amanzanero/advent-of-code-2020/lib"
	"regexp"
	"strconv"
)

type Op interface {
	Execute(vm *VMachine) error
	OpType() string
	Value() int
}

type AccOp struct {
	opType string
	value  int
}

func (op *AccOp) OpType() string {
	return op.opType
}

func (op *AccOp) Value() int {
	return op.value
}

func (op *AccOp) Execute(vm *VMachine) error {
	err := vm.AddHistory()
	if err != nil {
		return err
	}

	vm.acc += op.value
	vm.instruction += 1
	return nil
}

type JmpOp struct {
	opType string
	value  int
}

func (op *JmpOp) OpType() string {
	return op.opType
}

func (op *JmpOp) Value() int {
	return op.value
}

func (op *JmpOp) Execute(vm *VMachine) error {
	err := vm.AddHistory()
	if err != nil {
		return err
	}

	vm.instruction += op.value
	return nil
}

type NOp struct {
	opType string
	value  int
}

func (op *NOp) OpType() string {
	return op.opType
}

func (op *NOp) Value() int {
	return op.value
}

func (op *NOp) Execute(vm *VMachine) error {
	err := vm.AddHistory()
	if err != nil {
		return err
	}

	vm.instruction += 1
	return nil
}

func ParseOpFromString(line string) Op {
	reg := regexp.MustCompile(" ")
	splitLine := reg.Split(line, -1)

	opName := splitLine[0]
	value, err := strconv.Atoi(splitLine[1])
	lib.Check(err)

	var op Op
	switch opName {
	case "acc":
		op = &AccOp{opType: opName, value: value}
		break
	case "jmp":
		op = &JmpOp{opType: opName, value: value}
		break
	case "nop":
		op = &NOp{opName, value}
		break
	default:
		op = &NOp{"nop", 0}
		break
	}
	return op
}
