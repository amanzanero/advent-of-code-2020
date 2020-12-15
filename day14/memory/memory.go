package memory

type MemBlock map[uint64]uint64

type MemExpression struct {
	Value, Mask, MaskValue, Wildcard, Destination uint64
}

func (exp *MemExpression) wildcardBitLocations() []int {
	bitLocations := make([]int, 0)
	for i := 0; i < 36; i++ {
		var shift uint64 = 1 << i
		if shift == exp.Wildcard&shift {
			bitLocations = append(bitLocations, i)
		}
	}
	return bitLocations
}

func (exp *MemExpression) WildcardCombinations() []uint64 {
	bitLocations := exp.wildcardBitLocations()
	combinations := make([]uint64, 0)

	var combos func(int, uint64)
	combos = func(i int, sum uint64) {
		if i >= len(bitLocations) {
			combinations = append(combinations, sum)
			return
		}
		combos(i+1, sum+(1<<bitLocations[i]))
		combos(i+1, sum)
	}
	combos(0, 0)
	return combinations
}
