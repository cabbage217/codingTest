package car

import (
	mainConsts "../consts"
	"strings"
)

//car第一版
type CarV1 struct {
	orientation mainConsts.Direction
	x, y int
}

//创建并初始化car实例
func NewCarV1(orientation mainConsts.Direction) *CarV1 {
	return &CarV1{
		orientation: orientation,
		x: 1,
		y: 1,
	}
}

//执行动作，包括移动和转向
func (c *CarV1) Move(command string) {
	if command = strings.TrimSpace(command); command == "" {
		return
	}
	switch command {
	case "forward":
		switch c.orientation {
		case mainConsts.DirectionNorth:
			c.y += 1
		case mainConsts.DirectionSouth:
			c.y -= 1
		case mainConsts.DirectionWest:
			c.x -= 1
		case mainConsts.DirectionEast:
			c.x += 1
		default:
		}
	case "turn":
		c.orientation = (c.orientation + 1) % mainConsts.DirectionCount
	default:
		return
	}
}

//当前位置x坐标
func (c *CarV1) GetPositionX() int {
	return c.x
}

//当前位置y坐标
func (c *CarV1) GetPositionY() int {
	return c.y
}

////当前方向
func (c *CarV1) GetOrientation() string {
	return c.orientation.String()
}