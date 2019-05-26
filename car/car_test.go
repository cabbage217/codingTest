package car

import (
	"testing"
	"../consts"
	"fmt"
	"../command"
)

func TestCar(t *testing.T) {
	car := &Car{}
	car.Init(consts.DirectionNorth)
	commands := []carState{
		{command: command.ParseRestrainErr("f 1"), carStateBase: carStateBase{x: 1, y: 2, direction: consts.DirectionNorth}},
		{command: command.ParseRestrainErr("t"), carStateBase: carStateBase{x: 1, y: 2, direction: consts.DirectionEast}},
		{command: command.ParseRestrainErr("forward 1"), carStateBase: carStateBase{x: 2, y: 2, direction: consts.DirectionEast}},
		{command: command.ParseRestrainErr("turn"), carStateBase: carStateBase{x: 2, y: 2, direction: consts.DirectionSouth}},
	}
	doTestCar(t, car, commands)

	car.Init(consts.DirectionEast)
	commands = []carState{
		{command: command.ParseRestrainErr("f 1"), carStateBase: carStateBase{x: 2, y: 1, direction: consts.DirectionEast}},
		{command: command.ParseRestrainErr("t"), carStateBase: carStateBase{x: 2, y: 1, direction: consts.DirectionSouth}},
		{command: command.ParseRestrainErr("forward 1"), carStateBase: carStateBase{x: 2, y: 0, direction: consts.DirectionSouth}},
		{command: command.ParseRestrainErr("turn"), carStateBase: carStateBase{x: 2, y: 0, direction: consts.DirectionWest}},
	}
	doTestCar(t, car, commands)

	car.Init(consts.DirectionSouth)
	commands = []carState{
		{command: command.ParseRestrainErr("f 1"), carStateBase: carStateBase{x: 1, y: 0, direction: consts.DirectionSouth}},
		{command: command.ParseRestrainErr("t"), carStateBase: carStateBase{x: 1, y: 0, direction: consts.DirectionWest}},
		{command: command.ParseRestrainErr("forward 1"), carStateBase: carStateBase{x: 0, y: 0, direction: consts.DirectionWest}},
		{command: command.ParseRestrainErr("turn"), carStateBase: carStateBase{x: 0, y: 0, direction: consts.DirectionNorth}},
	}
	doTestCar(t, car, commands)

	car.Init(consts.DirectionWest)
	commands = []carState{
		{command: command.ParseRestrainErr("f 1"), carStateBase: carStateBase{x: 0, y: 1, direction: consts.DirectionWest}},
		{command: command.ParseRestrainErr("t"), carStateBase: carStateBase{x: 0, y: 1, direction: consts.DirectionNorth}},
		{command: command.ParseRestrainErr("forward 1"), carStateBase: carStateBase{x: 0, y: 2, direction: consts.DirectionNorth}},
		{command: command.ParseRestrainErr("turn"), carStateBase: carStateBase{x: 0, y: 2, direction: consts.DirectionEast}},
	}
	doTestCar(t, car, commands)
}

func doTestCar(t *testing.T, car *Car, commands []carState) {
	originalState := fmt.Sprintf("%v", car)
	for _, v := range commands {
		old := fmt.Sprintf("%v", car)
		car.Exec(v.command)
		if car.GetCurrentPosition().X != v.x || car.GetCurrentPosition().Y != v.y || car.GetOrientation().String() != v.direction.String() {
			t.Fatalf("command: %s, car state(%v) should be %s, old state: %s, original: %s", v.command, car, v.carStateBase.String(), old, originalState)
		}
	}
}