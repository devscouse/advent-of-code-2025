/*Package day4*/
package day4

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/devscouse/advent-of-code-2025/common"
)

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

type BitArray struct {
	bits []uint64
}

func NewBitArray(size int) *BitArray {
	numWords := (size + 63) / 64
	return &BitArray{bits: make([]uint64, numWords)}
}

func (b *BitArray) Set(idx int) {
	word := idx / 64
	wordIdx := idx % 64
	if word >= len(b.bits) {
		log.Printf("Expanding bitArray to %d words\n", word+1)
		newBits := make([]uint64, word+1)
		copy(newBits, b.bits)
		b.bits = newBits
	}
	b.bits[word] |= (1 << wordIdx)
}

func (b *BitArray) IsSet(idx int) bool {
	word := idx / 64
	wordIdx := idx % 64
	if word >= len(b.bits) {
		panic(errors.New("bounds exceeded on bitMap"))
	}
	return b.bits[word]&(1<<wordIdx) != 0
}

type BitMap struct {
	array  *BitArray
	width  int
	height int
}

func NewBitMap(bitArray *BitArray, width int, height int) *BitMap {
	return &BitMap{
		array:  bitArray,
		width:  width,
		height: height,
	}
}

func (b *BitMap) String() string {
	bitMapRepr := fmt.Sprintf("BitMap{\n   width: %d\n  height: %d\n\n", b.width, b.height)
	for y := range b.height {
		bitMapRepr += "  "
		for x := range b.width {
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
	bitIdx := y*b.width + x
	b.array.Set(bitIdx)
}

func (b *BitMap) IsSet(x int, y int) bool {
	bitIdx := y*b.width + x
	return b.array.IsSet(bitIdx)
}

func (b *BitMap) CountSurroundingSet(x int, y int) uint8 {
	count := uint8(0)
	for _, offset := range offsets {
		dx := x + offset[0]
		dy := y + offset[1]
		if dx < 0 || dx >= b.width {
			continue
		}
		if dy < 0 || dy >= b.height {
			continue
		}
		if b.IsSet(dx, dy) {
			count++
		}
	}
	return count
}

func ReadMap(bfr *bufio.Reader) *BitMap {
	currPos := 0
	mapWidth := 0
	mapHeight := 0
	bitArray := NewBitArray(64 * 500)

	for {
		r, _, err := bfr.ReadRune()
		if err == io.EOF {
			break
		}
		fmt.Printf("%c", r)
		switch r {
		case '\n':
			if mapWidth == 0 {
				mapWidth = currPos
			}
			mapHeight++
		case '.':
			currPos++
		case '@':
			bitArray.Set(currPos)
			currPos++
		default:
			log.Printf("Unrecognized rune: %c\n", r)
		}
	}
	return NewBitMap(bitArray, mapWidth, mapHeight)
}

func PartOne() {
	file := common.ReadPackageData("day4", "input.dat")
	bfr := bufio.NewReader(file)
	bitMap := ReadMap(bfr)
	count := 0

	for y := range bitMap.height {
		for x := range bitMap.height {
			if !bitMap.IsSet(x, y) {
				continue
			}
			if bitMap.CountSurroundingSet(x, y) < 4 {
				count++
			}
		}
	}
	fmt.Printf("Day 4 - Part One: %d\n", count)
}
