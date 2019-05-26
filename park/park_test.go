package park

import (
	"testing"
	"../position"
	"strings"
	"math/rand"
	"time"
	"fmt"
	"math"
)

func TestPark(t *testing.T) {
	p := &Park{}
	if _, err := p.IsPositionValid(position.Position{X: 2, Y: 3}); err == nil {
		t.Error("park has not initialized, function IsPositionValid shouldn't return a nil error\n")
	} else if !strings.Contains(err.Error(), "Park has not initialized") {
		t.Error("park has not initialized, function IsPositionValid should return an error with string \"Park has not initialized\"\n")
	}

	rand.Seed(time.Now().UnixNano())

	row := rand.Int()
	col := rand.Int()
	fmt.Printf("row: %d, col: %d\n", row, col)
	if err := p.Init(row, col); err == nil {
		if row <= 0 || col <= 0 {
			t.Errorf("park initialized with row(%d) and col(%d) should return an error\n", row, col)
		}
	} else if (row <= 0 || col <= 0) && !strings.Contains(err.Error(), "Invalid params") {
		t.Errorf("park initialized with row(%d) and col(%d) should return an error with string\"Invalid params\"\n", row, col)
	}

	count := rand.Intn(math.MaxUint16)
	fmt.Printf("for loop count: %d\n", count)
	for count > 0 {
		count--
		pos := position.Position{X: rand.Int(), Y: rand.Int()}
		//fmt.Printf("one position: %v\n", pos)
		if ok, err := p.IsPositionValid(pos); err == nil {
			if ok {
				if pos.X < 1 || pos.Y < 1 || pos.X > row || pos.Y > col {
					t.Error("function IsPositionValid shouldn't return true\n")
				}
			} else {
				if pos.X > 0 && pos.Y > 1 && pos.X <= row && pos.Y <= col {
					t.Error("function IsPositionValid shouldn't return false\n")
				}
			}
		} else {
			t.Errorf("function IsPositionValid shouldn't return an error(%s)\n", err.Error())
		}
		time.Sleep(time.Millisecond * 3)
	}
}
