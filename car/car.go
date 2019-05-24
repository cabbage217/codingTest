package car

import (
	"../command"
	"../position"
	"errors"
	"fmt"
	"../consts"
)

//第二版实现
type Car struct {
	orientation     consts.Direction
	currentPosition *position.Position
}

//初始化
func (c *Car) Init() {
	c.orientation = consts.DirectionNorth
	c.currentPosition = &position.Position{
		X: 1,
		Y: 1,
	}
}

//执行命令
func (c *Car) Exec(cmd *command.Command) error {
	if cmd == nil {
		return errors.New("Command is nil ")
	}
	switch cmd.Type {
	case command.TypeForward:
		return c.forward(cmd)
	case command.TypeTurn:
		return c.turn(cmd)
	default:
		return fmt.Errorf("Invalid command, %v ", cmd)
	}
}

//方向
func (c *Car) GetOrientation() consts.Direction {
	return c.orientation
}

//当前位置
func (c *Car) GetCurrentPosition() position.Position {
	return position.Position{
		X: c.currentPosition.X,
		Y: c.currentPosition.Y,
	}
}

//重置
func (c *Car) Reset() {
	c.Init()
}

//方便输出
func (c Car) String() string {
	return fmt.Sprintf("orientation: %s, currentPosition: %v", c.orientation, c.currentPosition)
}

//内部方法，移到
func (c *Car) forward(cmd *command.Command) error {
	switch c.orientation {
	case consts.DirectionNorth:
		c.currentPosition.Y += cmd.StepCount
	case consts.DirectionSouth:
		c.currentPosition.Y -= cmd.StepCount
	case consts.DirectionWest:
		c.currentPosition.X -= cmd.StepCount
	case consts.DirectionEast:
		c.currentPosition.X += cmd.StepCount
	default:
		return fmt.Errorf("Invalid direction %d ", cmd.Direction)
	}
	return nil
}

//内部方法，转向
func (c *Car) turn(cmd *command.Command) error {
	c.orientation = (c.orientation + 1) % consts.DirectionCount
	return nil
}
