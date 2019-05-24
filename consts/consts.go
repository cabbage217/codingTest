package consts

//定义方向类型，增强可读性
type Direction uint8

//方便输出
func (d Direction) String() string {
	switch d {
	case DirectionNorth:
		return "north"
	case DirectionSouth:
		return "south"
	case DirectionWest:
		return "west"
	case DirectionEast:
		return "east"
	default:
		return "unknown"
	}
}

//方向枚举类型总个数
const DirectionCount = 4

//方向枚举类型个数
const (
	DirectionNorth Direction = 0
	DirectionEast  Direction = 1
	DirectionSouth Direction = 2
	DirectionWest  Direction = 3
)
