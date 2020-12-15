package utils

import (
	"github.com/amanzanero/advent-of-code-2020/day14/memory"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"regexp"
	"strconv"
)

func ParseMask(maskExpression string) (mask uint64, value uint64, wildcard uint64) {
	reg := regexp.MustCompile(" = ")
	splitInput := reg.Split(maskExpression, 2)
	maskInput := splitInput[1]

	mask = 0
	value = 0
	wildcard = 0
	for i := len(maskInput) - 1; i >= 0; i-- {
		shift := len(maskInput) - 1 - i
		switch maskInput[i] {
		case 'X':
			wildcard |= 1 << shift
			continue
		case '0':
			mask |= 1 << shift
		case '1':
			mask |= 1 << shift
			value |= 1 << shift
		}
	}
	return
}

func ParseExpression(input string) *memory.MemExpression {
	reg := regexp.MustCompile("] = ")
	splitInput := reg.Split(input, 2)
	expressionInput := splitInput[1]

	expression, err := strconv.ParseUint(expressionInput, 10, 64)
	lib.Check(err)

	destinationInput := splitInput[0][4:]
	destination, destErr := strconv.ParseUint(destinationInput, 10, 64)
	lib.Check(destErr)

	return &memory.MemExpression{Destination: destination, Value: expression}
}
