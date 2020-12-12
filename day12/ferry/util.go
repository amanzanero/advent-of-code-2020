package ferry

func rotate(curr, degrees int, left bool) int {
	var newDegrees int
	if left {
		newDegrees = (curr - degrees + 360) % 360
	} else {
		newDegrees = (curr + degrees) % 360
	}
	return newDegrees
}

func getEndPosition(x, y, unit int, direction string) (endX int, endY int) {
	delta := DIRECTIONS[direction]
	endX = x + delta[0]*unit
	endY = y + delta[1]*unit
	return
}
