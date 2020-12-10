package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"sort"
)

func main() {
	parsedNums := lib.GetIntArrayInput("day10/input.txt")
	sort.Ints(parsedNums)

	partOne(parsedNums)
	partTwoCachedRecursive(parsedNums)
}

func partOne(parsedNums []int) {
	defer lib.Elapsed("--took:")()
	diff, prev, ones, threes := 0, 0, 0, 1

	for _, num := range parsedNums {
		diff = num - prev
		prev = num

		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		} else if diff > 3 {
			break
		}

	}
	fmt.Printf("part one: %d\n", ones*threes)
}

func partTwoCachedRecursive(parsedNums []int) {
	defer lib.Elapsed("--took:")()
	cache := make([]int, len(parsedNums))

	var dfs func(int) int
	dfs = func(curr int) int {
		if curr == len(parsedNums)-1 {
			cache[curr] = 1
			return 1
		}

		if cache[curr] > 0 {
			return cache[curr]
		}

		resultsFromHere := 0
		for next := curr + 1; next < len(parsedNums); next++ {
			if parsedNums[next]-parsedNums[curr] <= 3 {
				resultsFromHere += dfs(next)
			} else {
				break
			}
		}
		cache[curr] = resultsFromHere
		return resultsFromHere
	}
	ans := 0
	for i, num := range parsedNums {
		if num <= 3 {
			ans += dfs(i)
		}
	}
	fmt.Printf("part two: %d\n", ans)
}
