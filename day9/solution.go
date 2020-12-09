package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/day9/nqueue"
	"github.com/amanzanero/advent-of-code-2020/day9/utils"
	"github.com/amanzanero/advent-of-code-2020/lib"
	"math"
)

func main() {
	parsedNums := lib.GetIntArrayInput("day9/input.txt")

	pOneAns := partOne(parsedNums, 25)
	fmt.Println()
	partTwo(pOneAns, parsedNums)
}

func partOne(nums []int, preamble int) int {
	defer lib.Elapsed("--time elapsed:")()

	ans := -1
	// aggregate nqueue
	nq := nqueue.NewNQueue(preamble)
	for i, numA := range nums[:preamble] {
		sums := make([]int, 0)
		for _, numB := range nums[0:i] {
			sums = append(sums, numA+numB)
		}
		nq.PushPop(sums)
	}

	// now find the first bad input
	for i, num := range nums[preamble:] {
		if !nq.Has(num) {
			ans = num
			break
		}
		sums := make([]int, 0)
		for _, prevNum := range nums[i+1 : i+preamble] {
			sums = append(sums, num+prevNum)
		}
		nq.PushPop(sums)
	}
	fmt.Printf("Day 9 part 1: %d\n", ans)
	return ans
}

func partTwo(partOneAnswer int, nums []int) {
	defer lib.Elapsed("--time elapsed:")()
	ans := -1

	memo := utils.NbyNSlice(len(nums))
	// fill out base cases
	for i := range memo {
		memo[i][i] = nums[i]
	}

	for n := 1; n < len(nums); n++ {
		for i := 0; i < len(nums)-n; i++ {
			j := i + n
			res := memo[i][j-1] + memo[j][j]
			if res == partOneAnswer { // found the range!
				min, max := math.MaxFloat64, -math.MaxFloat64
				for _, num := range nums[i : j+1] {
					min = math.Min(min, float64(num))
					max = math.Max(max, float64(num))
				}
				ans = int(min + max)
				fmt.Printf("Day 9 part 2: %d\n", ans)
				return
			}
			memo[i][j] = res
		}

	}
}
