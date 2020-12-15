package main

import "fmt"

//var input = []int{0, 3, 6}
var input = []int{14, 1, 17, 0, 3, 20}

//var ITERATION = 2020
var ITERATION = 30000000

func main() {
	spokenWords := make(map[int][]int) // number: last spoken

	for i, num := range input {
		seen := make([]int, 1)
		seen[0] = i
		spokenWords[num] = seen
	}

	spoken := input[len(input)-1]
	for i := len(input); i < ITERATION; i++ {
		lastTime, exists := spokenWords[spoken]
		if exists {
			if len(lastTime) == 1 {
				spoken = 0
			} else {
				spoken = lastTime[1] - lastTime[0]
			}
		}

		currSeen, currSeenExists := spokenWords[spoken]
		if currSeenExists {
			if len(currSeen) == 1 {
				spokenWords[spoken] = append(spokenWords[spoken], i)
			} else {
				currSeen[0], currSeen[1] = currSeen[1], i
			}
		} else {
			seen := make([]int, 1)
			seen[0] = i
			spokenWords[spoken] = seen
		}
	}
	fmt.Println(spoken)
}
