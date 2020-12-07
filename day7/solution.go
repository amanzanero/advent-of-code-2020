package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/day7/rules"
	"github.com/amanzanero/advent-of-code-2020/lib"
)

func main() {
	parsedRules := lib.ParseLines("day7/input.txt")
	ruleMap := buildRuleMap(parsedRules)

	fmt.Printf("Day 7 part 1: %d\n", partOne(ruleMap))
	fmt.Printf("Day 7 part 2: %d\n", partTwo(ruleMap))
}

func partOne(ruleMap map[string]*rules.Rule) int {
	// do a bfs on each node/rule starting from shiny gold bag
	root := ruleMap["shiny gold"]
	queue := make([]*rules.Rule, 0)
	queue = append(queue, root)

	colors := make(map[string]bool)

	firstBag := true
	for len(queue) != 0 {
		// pop
		curr := queue[0]
		queue = queue[1:]

		// increase the counter for number of potential outer bags
		if !firstBag {
			colors[curr.Color] = true
		} else {
			firstBag = false
		}

		for _, node := range curr.ContainedBy {
			queue = append(queue, node)
		}
	}

	return len(colors)
}

type TraversalNode struct {
	LevelCount int
	TRule      *rules.Rule
}

func partTwo(ruleMap map[string]*rules.Rule) int {
	root := ruleMap["shiny gold"]
	queue := make([]*TraversalNode, 0)
	queue = append(queue, &TraversalNode{1, root})

	numBags := 0

	firstBag := true
	for len(queue) != 0 {
		// pop
		curr := queue[0]
		queue = queue[1:]

		// increase the counter for number of potential outer bags
		if !firstBag {
			numBags += curr.LevelCount
		} else {
			firstBag = false
		}

		for color, count := range curr.TRule.Contains {
			rule := ruleMap[color]
			levelCount := curr.LevelCount * count
			queue = append(queue, &TraversalNode{levelCount, rule})
		}
	}

	return numBags
}

func buildRuleMap(lines []string) map[string]*rules.Rule {
	ruleMap := make(map[string]*rules.Rule)

	// step1. parse rules
	for _, line := range lines {
		rule := rules.ParseLineToRule(line)
		ruleMap[rule.Color] = rule
	}
	// step2. link rules by which other bags they can be contained by (parents)
	for _, rule := range ruleMap {
		for contains := range rule.Contains {
			containsRule := ruleMap[contains]
			containsRule.AddOuterBag(rule)
		}
	}
	return ruleMap
}
