package position

import "fmt"

//封装位置信息
type Position struct {
	X, Y int
}

//方便输出
func (p *Position) String() string {
	return fmt.Sprintf("{X: %d, Y: %d}", p.X, p.Y)
}
