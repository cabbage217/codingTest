package car

import (
	"../consts"
	"../command"
	"fmt"
)

type carStateBase struct {
	x, y int
	direction consts.Direction
}

func (c *carStateBase) String() string {
	return fmt.Sprintf("{x: %d, y: %d, direction: %s}", c.x, c.y, c.direction)
}

type carState struct {
	command *command.Command
	carStateBase
}
