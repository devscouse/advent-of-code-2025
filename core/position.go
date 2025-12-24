package core

type Pos struct {
	X int
	Y int
}

func NewPos(x int, y int) *Pos {
	return &Pos{X: x, Y: y}
}

func (p *Pos) Equals(other *Pos) bool {
	return p.X == other.X && p.Y == other.Y
}
