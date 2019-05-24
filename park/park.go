package park

import "fmt"
import (
	"../position"
	"errors"
)

//ParkInterface的实现
type Park struct {
	inited     bool
	xMax, yMax int
}

//初始化，未初始化状态时不可用，因为停车场长和宽未定义
func (p *Park) Init(row, col int) error {
	if row <= 0 || col <= 0 {
		return fmt.Errorf("Invalid params(row=%d, col=%d) where init a new park. ", row, col)
	}

	p.xMax = row
	p.yMax = col
	p.inited = true
	return nil
}

//重置初始化状态
func (p *Park) Reset() {
	p.xMax = 0
	p.yMax = 0
	p.inited = false
}

//检查是否已初始化
func (p *Park) IsInited() bool {
	return p.inited
}

//检查某个位置是否合法
func (p *Park) IsPositionValid(position position.Position) (bool, error) {
	if !p.IsInited() {
		return false, errors.New("Park has not inited ")
	}
	return position.X > 0 && position.X <= p.xMax && position.Y > 0 && position.Y <= p.yMax, nil
}

//停车场最大x值
func (p *Park) GetMaxX() int {
	return p.xMax
}

//停车场最大y值
func (p *Park) GetMaxY() int {
	return p.yMax
}

//方便输出
func (p *Park) String() string {
	return fmt.Sprintf("{inited: %t, xMax: %d, yMax: %d}", p.inited, p.xMax, p.yMax)
}
