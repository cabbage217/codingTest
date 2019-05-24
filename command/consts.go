package command

import "math"

//定义命令类型，增加可读性
type Type uint8

//方便输出
func (t Type) String() string {
	switch t {
	case TypeHelp:
		return "help"
	case TypeInitPark:
		return "initPark"
	case TypeForward:
		return "forward"
	case TypeTurn:
		return "turn"
	case TypeQuit:
		return "quit"
	default:
		return "unknown"
	}
}

const (
	TypeHelp     Type = 0
	TypeInitPark Type = 1
	TypeForward  Type = 2
	TypeTurn     Type = 3
	TypeQuit     Type = math.MaxUint8
)
