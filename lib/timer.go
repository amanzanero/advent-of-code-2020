package lib

import (
	"fmt"
	"time"
)

func Elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s %v\n", what, time.Since(start))
	}
}
