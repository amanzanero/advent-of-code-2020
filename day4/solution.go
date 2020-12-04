package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"regexp"
)

var PASSPORT_POLICIES = map[string]string{
	"byr": "(19[2-9][0-9]|200[0-2]){1}",
	"iyr": "(20(1[0-9]|20)){1}",
	"eyr": "(20(2[0-9]|30)){1}",
	"hgt": "(1[5-8][0-9]cm|19[0-3]cm|59in|6[0-9]in|7[0-6]in)",
	"hcl": "#[0-9a-f]{6}",
	"ecl": "(amb|blu|brn|gry|grn|hzl|oth){1}",
	"pid": "[0-9]{9}",
}

type AttrSet map[string]string
type PassportChecker func(passport AttrSet) bool

func main() {
	fmt.Println(countValidPassports(isValidPassport1))
	fmt.Println(countValidPassports(isValidPassport2))
}

func countValidPassports(checker PassportChecker) int {
	parsedLines := lib.ParseLines("day4/input.txt")
	count := 0
	iter := 0

	for iter < len(parsedLines) {
		// iterate until blank line or end
		passport := make(AttrSet)
		for _, line := range parsedLines[iter:] {
			iter++
			if line == "" {
				break
			}
			extracted := extractAttrs(line)
			passport = combineAttrs(passport, extracted)
		}

		if checker(passport) {
			count++
		}
	}
	return count
}

func extractAttrs(line string) AttrSet {
	attrs := make(AttrSet)
	delimiter := regexp.MustCompile(`\s+`)
	splitLine := delimiter.Split(line, -1)

	for _, attr := range splitLine {
		colonDelim := regexp.MustCompile(`:`)
		pair := colonDelim.Split(attr, -1)
		k, v := pair[0], pair[1]
		attrs[k] = v
	}
	return attrs
}

func combineAttrs(m1 AttrSet, m2 AttrSet) AttrSet {
	for k, v := range m2 {
		m1[k] = v
	}
	return m1
}

func isValidPassport1(passport AttrSet) bool {
	for attr, _ := range PASSPORT_POLICIES {
		_, exists := passport[attr]
		if !exists {
			return false
		}
	}
	return true
}

func isValidPassport2(passport AttrSet) bool {
	for attr, policy := range PASSPORT_POLICIES {
		value, exists := passport[attr]
		if !exists {
			return false
		}
		match, _ := regexp.MatchString(policy, value)
		if !match {
			return false
		}
	}
	return true
}
