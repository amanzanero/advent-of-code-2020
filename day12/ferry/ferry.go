package ferry

type ShipState struct {
	direction string
	position  []int
	waypoint  []int
}

func NewShipState() *ShipState {
	ss := ShipState{direction: "E"}
	ss.position = []int{0, 0}
	ss.waypoint = []int{1, 10}
	return &ss
}

func (ss *ShipState) CalculateManhattan() int {
	dx, dy := ss.position[0], ss.position[1]
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}
	return dy + dx
}
