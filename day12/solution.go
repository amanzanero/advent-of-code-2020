package main

import (
	"fmt"
	"github.com/amanzanero/advent-of-code-2020/day12/ferry"
	"github.com/amanzanero/advent-of-code-2020/lib"
)

func main() {
	parsedInputs := lib.ParseLines("day12/input.txt")
	navigationInstructions := make([]ferry.Navigation, 0)
	for _, in := range parsedInputs {
		nav, err := ferry.ParseCommand(in)
		if err != nil {
			fmt.Println(err)
		}
		navigationInstructions = append(
			navigationInstructions,
			nav,
		)
	}
	partOne(navigationInstructions)
	partTwo(navigationInstructions)
}

func partOne(navigations []ferry.Navigation) {
	defer lib.Elapsed("-- took")()
	ss := ferry.NewShipState()

	for _, nav := range navigations {
		nav.MovePosition(ss)
	}

	fmt.Printf("Sum of manhattan: %d\n", ss.CalculateManhattan())
}

func partTwo(navigations []ferry.Navigation) {
	defer lib.Elapsed("-- took")()
	ss := ferry.NewShipState()

	for _, nav := range navigations {
		nav.MoveWaypoint(ss)
	}

	fmt.Printf("Sum of manhattan: %d\n", ss.CalculateManhattan())
}
