package park

import (
	"../position"
)

type ParkInterface interface {
	Init(row, col int) error
	Reset()
	IsInitialized() bool
	IsPositionValid(position position.Position) (bool, error)
	GetMaxX() int
	GetMaxY() int
}
