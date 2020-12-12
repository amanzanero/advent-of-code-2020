package ferry

import (
	"errors"
	"fmt"
	"strconv"
)

func ParseCommand(in string) (Navigation, error) {
	command := string(in[0])
	unit, err := strconv.Atoi(in[1:])

	if err != nil {
		return nil, err
	}

	switch command {
	case "N", "S", "E", "W":
		return &DirectionAction{
			direction: command,
			unit:      unit,
		}, nil
	case "L", "R":
		return &RotateAction{
			left:    command == "L",
			degrees: unit,
		}, nil
	case "F":
		return &ForwardAction{unit: unit}, nil
	default:
		return nil, errors.New(fmt.Sprintf("unrecognized command: %s", command))
	}
}
