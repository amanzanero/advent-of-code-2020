package ferry

var DIRECTIONS = map[string][]int{
	"N": {1, 0},
	"S": {-1, 0},
	"E": {0, 1},
	"W": {0, -1},
}

var DEGREES2DIRECTION = map[int]string{
	0:   "E",
	90:  "S",
	180: "W",
	270: "N",
}

var DIRECTION2DEGREES = map[string]int{
	"E": 0,
	"S": 90,
	"W": 180,
	"N": 270,
}

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

type Navigation interface {
	MovePosition(ss *ShipState)
	MoveWaypoint(ss *ShipState)
}

type DirectionAction struct {
	direction string
	unit      int
}

func (da *DirectionAction) MovePosition(ss *ShipState) {
	newX, newY := getEndPosition(ss.position[0], ss.position[1], da.unit, da.direction)
	ss.position[0], ss.position[1] = newX, newY
}

func (da *DirectionAction) MoveWaypoint(ss *ShipState) {
	newX, newY := getEndPosition(ss.waypoint[0], ss.waypoint[1], da.unit, da.direction)
	ss.waypoint[0], ss.waypoint[1] = newX, newY
}

type RotateAction struct {
	left    bool
	degrees int
}

func (ra *RotateAction) MovePosition(ss *ShipState) {
	newDegrees := rotate(DIRECTION2DEGREES[ss.direction], ra.degrees, ra.left)
	ss.direction = DEGREES2DIRECTION[newDegrees]
}

func (ra *RotateAction) MoveWaypoint(ss *ShipState) {
	rotations := ra.degrees / 90
	for i := 0; i < rotations; i++ {
		if ra.left {
			ss.waypoint[0], ss.waypoint[1] = ss.waypoint[1], -ss.waypoint[0]
		} else {
			ss.waypoint[0], ss.waypoint[1] = -ss.waypoint[1], ss.waypoint[0]
		}
	}
}

type ForwardAction struct {
	unit int
}

func (fa *ForwardAction) MovePosition(ss *ShipState) {
	newX, newY := getEndPosition(ss.position[0], ss.position[1], fa.unit, ss.direction)
	ss.position[0], ss.position[1] = newX, newY
}

func (fa *ForwardAction) MoveWaypoint(ss *ShipState) {
	ss.position[0] += ss.waypoint[0] * fa.unit
	ss.position[1] += ss.waypoint[1] * fa.unit
}

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
