package rules

import (
	"github.com/amanzanero/advent-of-code-2020/lib"
	"regexp"
	"strconv"
	"strings"
)

// Analogous to a node
type Rule struct {
	Color       string         // val
	Contains    map[string]int // children
	ContainedBy []*Rule        // parents
}

func newRule(color string) *Rule {
	return &Rule{Color: color}
}

func ParseLineToRule(line string) *Rule {
	split := splitString(line, ` bags contain `)

	color := split[0]
	rule := newRule(color)
	rule.ContainedBy = make([]*Rule, 0)
	rule.Contains = make(map[string]int)

	trimmedRules := strings.Trim(split[1], ".")
	if trimmedRules == "no other bags" {
		return rule
	}

	splitRules := splitString(trimmedRules, `, `)
	for _, parsedRule := range splitRules {
		countColorList := splitStringNTimes(parsedRule, ` `, 2)

		count, err := strconv.Atoi(countColorList[0])
		lib.Check(err)

		countColor := trimBagsFromColor(countColorList[1], count)
		rule.Contains[countColor] = count
	}

	return rule
}

func (r *Rule) AddOuterBag(rule *Rule) {
	r.ContainedBy = append(r.ContainedBy, rule)
}

func splitString(line, delim string) []string {
	reg := regexp.MustCompile(delim)
	split := reg.Split(line, -1)
	return split
}

func splitStringNTimes(line, delim string, n int) []string {
	reg := regexp.MustCompile(delim)
	split := reg.Split(line, n)
	return split
}

func trimBagsFromColor(color string, count int) string {
	var end int
	if count == 1 {
		end = 4
	} else {
		end = 5
	}
	return color[:len(color)-end]
}
