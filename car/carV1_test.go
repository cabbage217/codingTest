package car

import (
	"testing"
	"../consts"
	"fmt"
)

func TestCarV1(t *testing.T) {
	keys := []string {"f", "t", "forward", "turn"}
	car := NewCarV1(consts.DirectionNorth)
	commandStrings := map[string]carStateBase{
		"f": {x: 1, y: 2, direction: consts.DirectionNorth},
		"t": {x: 1, y: 2, direction: consts.DirectionEast},
		"forward": {x: 2, y: 2, direction: consts.DirectionEast},
		"turn": {x: 2, y: 2, direction: consts.DirectionSouth},
	}
	doTestCarV1(t, car, commandStrings, keys)

	car = NewCarV1(consts.DirectionEast)
	commandStrings = map[string]carStateBase{
		"f": {x: 2, y: 1, direction: consts.DirectionEast},
		"t": {x: 2, y: 1, direction: consts.DirectionSouth},
		"forward": {x: 2, y: 0, direction: consts.DirectionSouth},
		"turn": {x: 2, y: 0, direction: consts.DirectionWest},
	}
	doTestCarV1(t, car, commandStrings, keys)

	car = NewCarV1(consts.DirectionSouth)
	commandStrings = map[string]carStateBase{
		"f": {x: 1, y: 0, direction: consts.DirectionSouth},
		"t": {x: 1, y: 0, direction: consts.DirectionWest},
		"forward": {x: 0, y: 0, direction: consts.DirectionWest},
		"turn": {x: 0, y: 0, direction: consts.DirectionNorth},
	}
	doTestCarV1(t, car, commandStrings, keys)

	car = NewCarV1(consts.DirectionWest)
	commandStrings = map[string]carStateBase{
		"f": {x: 0, y: 1, direction: consts.DirectionWest},
		"t": {x: 0, y: 1, direction: consts.DirectionNorth},
		"forward": {x: 0, y: 2, direction: consts.DirectionNorth},
		"turn": {x: 0, y: 2, direction: consts.DirectionEast},
	}
	doTestCarV1(t, car, commandStrings, keys)
}

func doTestCarV1(t *testing.T, car *CarV1, commandStrings map[string]carStateBase, keys []string) {
	originalState := fmt.Sprintf("%v", car)
	for _, k := range keys {
		v := commandStrings[k]
		old := fmt.Sprintf("%v", car)
		car.Move(k)
		if car.GetPositionX() != v.x || car.GetPositionY() != v.y || car.GetOrientation() != v.direction.String() {
			t.Fatalf("command: %s, car state(%v) should be %s, old state: %s, original: %s", k, car, v.String(), old, originalState)
		}
	}
}
