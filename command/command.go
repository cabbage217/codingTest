package command

import (
	mainConsts "../consts"
	"fmt"
	"strconv"
	"strings"
)

//用户输入对应的命令
var commandsMap map[string]Type

//初始化用户输入对应的命令，只会执行一次，并且多线程安全
func init() {
	commandsMap = map[string]Type{
		"p":       TypeInitPark,
		"park":    TypeInitPark,
		"f":       TypeForward,
		"forward": TypeForward,
		"t":       TypeTurn,
		"turn":    TypeTurn,
		"q":       TypeQuit,
		"quit":    TypeQuit,
	}
}

//封闭命令
type Command struct {
	Type Type
	Direction   mainConsts.Direction
	X, Y        int
	StepCount   int
}

//打印命令帮助手册
func PrintHelp() {
	fmt.Print(`
Commands Available:
p (or park)    - Init park map with max x and max y, example: p 4 4 or park 4 4.
f (or forward) - Forward, example f 1 or forward 1
t (or turn)    - Turn clockwise once, example t or turn
q (or quit)    - Quit, example q or quit
`)
}

//从用户输入解析命令
func Parse(commandLine string) (*Command, error) {
	if commandLine = strings.TrimSpace(commandLine); commandLine == "" {
		return nil, nil
	}

	inputs := strings.Split(commandLine, " ")
	inputCount := len(inputs)
	if inputCount < 1 {
		return nil, nil
	}

	if commandType, ok := commandsMap[inputs[0]]; ok {
		switch commandType {
		case TypeQuit:
			return &Command{Type: TypeQuit}, nil
		case TypeInitPark:
			if inputCount < 3 {
				return nil, fmt.Errorf("Invalid command, params no enough ")
			}
			x, e := strconv.Atoi(inputs[1])
			if e != nil {
				return nil, fmt.Errorf("Can not parse param x(%s) of command park ", inputs[1])
			}
			y, e := strconv.Atoi(inputs[2])
			if e != nil {
				return nil, fmt.Errorf("Can not parse param y(%s) of command park ", inputs[2])
			}
			c := &Command{
				Type: TypeInitPark,
				X:           x,
				Y:           y,
			}
			return c, nil
		case TypeForward:
			if inputCount < 2 {
				return nil, fmt.Errorf("Invalid command, params no enough ")
			}
			step, e := strconv.Atoi(inputs[1])
			if e != nil {
				return nil, fmt.Errorf("Can not parse param step(%s) of command forward ", inputs[1])
			}
			return &Command{Type: TypeForward, StepCount: step}, nil
		case TypeTurn:
			return &Command{Type: TypeTurn}, nil
		case TypeHelp:
			return &Command{Type: TypeHelp}, nil
		default:
			return nil, fmt.Errorf("Unknown command \"%s\" ", commandLine)
		}
	} else {
		return nil, fmt.Errorf("Invalid command \"%s\" ", commandLine)
	}

	return nil, nil
}

//方便输出
func (c *Command) String() string {
	return fmt.Sprintf("{Type: %v, Direction: %v, X: %d, Y: %d, StepCount: %d}",
		c.Type, c.Direction, c.X, c.Y, c.StepCount)
}
