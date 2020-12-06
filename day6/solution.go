package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/lib"
)

type Answers map[uint8]bool
type SetOperation func(a, b Answers) Answers

func main() {
	parsedInput := lib.ParseLines("day6/input.txt")

	fmt.Printf("Part 1: %d\n", partOne(parsedInput))
	fmt.Printf("Part 2: %d\n", partTwo(parsedInput))
}

func partOne(input []string) int {
	return sumTotalsOfGroupAnswers(input, union)
}

func partTwo(input []string) int {
	return sumTotalsOfGroupAnswers(input, intersection)
}

func sumTotalsOfGroupAnswers(input []string, setOp SetOperation) int {
	totals := make([]int, 0)

	var groupAnswers Answers = nil
	for _, line := range input {
		if line == "" {
			totals = append(totals, len(groupAnswers))
			groupAnswers = nil
			continue
		}

		singlePersonAnswer := getAnswersFromLine(line)
		if groupAnswers != nil {
			groupAnswers = setOp(groupAnswers, singlePersonAnswer)
		} else {
			groupAnswers = singlePersonAnswer
		}
	}
	totals = append(totals, len(groupAnswers))

	sum := 0
	for _, total := range totals {
		sum += total
	}
	return sum
}

func getAnswersFromLine(line string) Answers {
	answers := make(Answers)
	for i := 0; i < len(line); i++ {
		answers[line[i]] = true
	}
	return answers
}

func union(a, b Answers) Answers {
	for k, v := range b {
		a[k] = v
	}
	return a
}

func intersection(a, b Answers) Answers {
	resultSet := make(Answers)

	for key := range a {
		if b[key] {
			resultSet[key] = true
		}
	}

	for key := range b {
		if a[key] {
			resultSet[key] = true
		}
	}

	return resultSet
}
