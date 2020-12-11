package ferry

var DELTAS = [][]int{
	{-1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
	{1, 0},
	{1, -1},
	{0, -1},
	{-1, -1},
}

const SEAT uint8 = '#'
const EMPTY uint8 = 'L'
const FLOOR uint8 = '.'

func NbyMMatrix(n, m int) [][]uint8 {
	matrix := make([][]uint8, n)
	for i := range matrix {
		matrix[i] = make([]uint8, m)
	}
	return matrix
}

func isValidDelta(x, y, nRows, nCols int) bool {
	return x >= 0 && x < nRows && y >= 0 && y < nCols
}

func PersonCanSit(x, y int, rows [][]uint8) bool {
	for _, delta := range DELTAS {
		dx, dy := delta[0], delta[1]
		if !isValidDelta(x+dx, y+dy, len(rows), len(rows[0])) {
			continue
		}

		if rows[x+dx][y+dy] == SEAT {
			return false
		}
	}
	return true
}

func PersonCanSitFirstEachDirection(x, y int, rows [][]uint8) bool {
	for _, delta := range DELTAS {
		dx, dy := delta[0], delta[1]
		validateX, validateY := x+dx, y+dy
		for isValidDelta(validateX, validateY, len(rows), len(rows[0])) {
			if rows[validateX][validateY] == SEAT {
				return false
			} else if rows[validateX][validateY] == EMPTY {
				break
			}
			validateX, validateY = validateX+dx, validateY+dy
		}
	}
	return true
}

func PersonWillLeave(x, y int, rows [][]uint8) bool {
	count := 0
	for _, delta := range DELTAS {
		dx, dy := delta[0], delta[1]
		if !isValidDelta(x+dx, y+dy, len(rows), len(rows[0])) {
			continue
		}

		if rows[x+dx][y+dy] == SEAT {
			count++
		}
	}
	return count >= 4
}

func TolerantPersonWillLeave(x, y int, rows [][]uint8) bool {
	count := 0
	for _, delta := range DELTAS {
		dx, dy := delta[0], delta[1]
		validateX, validateY := x+dx, y+dy
		for isValidDelta(validateX, validateY, len(rows), len(rows[0])) {
			if rows[validateX][validateY] == SEAT {
				count += 1
				break
			} else if rows[validateX][validateY] == EMPTY {
				break
			}
			validateX, validateY = validateX+dx, validateY+dy
		}
	}
	return count >= 5
}
