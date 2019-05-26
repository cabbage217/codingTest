package command

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	m := map[string]Type{
		"p 5 6":     TypeInitPark,
		"park 5 6":  TypeInitPark,
		"f 2":       TypeForward,
		"forward 2": TypeForward,
		"t":         TypeTurn,
		"turn":      TypeTurn,
		"q":         TypeQuit,
		"quit":      TypeQuit,
		"h":         TypeHelp,
		"help":      TypeHelp,
	}

	for k, v := range m {
		fmt.Printf("k: %s, v: %d\n", k, v)
		if command, err := Parse(k); err != nil {
			t.Errorf("Failed to parse command: %s, error msg: %s\n", k, err.Error())
		} else if command == nil {
			t.Errorf("Failed to parse command: %s, result shouldn't be nil\n", k)
		} else if command.Type != v {
			t.Errorf("Failed to parse command: %s, result type(%d) not correct, should be %d\n", k, command.Type, v)
		}
	}

	rand.Seed(time.Now().UnixNano())
	count := rand.Intn(math.MaxUint8)
	fmt.Println("test count: ", count)
	validCommands := map[string]uint8{
		"p": 2,
		"f": 1,
		"t": 0,
		"q": 0,
		"h": 0,
	}
	for count > 0 {
		count--
		commandStr := randomChar()
		var (
			paramCount uint8
			isValid    bool
			params     = make([]interface{}, 0)
			isParamsValid = true
		)
		if paramCount, isValid = validCommands[commandStr]; isValid {
			for paramCount > 0 {
				paramCount--
				oneParam := rand.Int()
				params = append(params, oneParam)
				commandStr += " " + strconv.Itoa(oneParam)
			}
			if paramCount > 0 {
				for _, v := range params {
					if v.(int) <= 0 {
						isParamsValid = false
						break
					}
				}
			}
		}
		fmt.Println("rand command string: ", commandStr)
		if _, err := Parse(commandStr); err != nil {
			if isValid && isParamsValid {
				t.Errorf("Failed to parse command: %s, error msg: %s\n", commandStr, err.Error())
			}
		} else {
			if !isValid || !isParamsValid {
				t.Errorf("Failed to test pares command: %s, should return an error\n", commandStr)
			}
		}
	}
}

func randomChar() string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	idx := rand.Intn(len(chars) - 1)
	return chars[idx : idx+1]
}
