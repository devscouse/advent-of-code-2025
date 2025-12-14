package common

import "fmt"

var offsets [8][2]int = [8][2]int{
	{1, 1},
	{1, 0},
	{1, -1},
	{0, 1},
	{0, -1},
	{-1, 1},
	{-1, 0},
	{-1, -1},
}

type BitMap struct {
	Array  *BitArray
	Width  int
	Height int
}

func NewBitMap(bitArray *BitArray, width int, height int) *BitMap {
	return &BitMap{
		Array:  bitArray,
		Width:  width,
		Height: height,
	}
}

func (b *BitMap) String() string {
	bitMapRepr := fmt.Sprintf("BitMap{\n   width: %d\n  height: %d\n\n", b.Width, b.Height)
	for y := range b.Height {
		bitMapRepr += "  "
		for x := range b.Width {
			if b.IsSet(x, y) {
				bitMapRepr += "@"
			} else {
				bitMapRepr += "."
			}
		}
		bitMapRepr += "\n"
	}
	bitMapRepr += "}"
	return bitMapRepr
}

func (b *BitMap) Set(x int, y int) {
	bitIdx := y*b.Width + x
	b.Array.Set(bitIdx)
}

func (b *BitMap) Unset(x int, y int) {
	b.Array.Unset(b.GetPositionIdx(x, y))
}

func (b *BitMap) IsSet(x int, y int) bool {
	bitIdx := y*b.Width + x
	return b.Array.IsSet(bitIdx)
}

func (b *BitMap) GetPositionIdx(x int, y int) int {
	return y*b.Width + x
}

func (b *BitMap) CountSurroundingSet(x int, y int) uint8 {
	count := uint8(0)
	for _, offset := range offsets {
		dx := x + offset[0]
		dy := y + offset[1]
		if dx < 0 || dx >= b.Width {
			continue
		}
		if dy < 0 || dy >= b.Height {
			continue
		}
		if b.IsSet(dx, dy) {
			count++
		}
	}
	return count
}
