// Package day7
package day7

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"slices"

	"github.com/devscouse/advent-of-code-2025/core"
)

const (
	startingBeamCapacity     = 100
	startingSplitterCapacity = 100
)

var ErrBeamLeftManifold error = errors.New("beam has left the bounds of the manifold")

type Manifold struct {
	width     int
	height    int
	splitters *core.BitMap
}

func NewManifold(splitters *core.BitMap) *Manifold {
	return &Manifold{width: splitters.Width, height: splitters.Height, splitters: splitters}
}

type TachyonManifold struct {
	manifold *Manifold
	beamY    int
	beamX    []int
	nSplits  int
}

func (m *TachyonManifold) String() string {
	repr := ""
	for y := range m.manifold.height {
		for x := range m.manifold.width {
			represented := false
			if m.manifold.splitters.IsSet(x, y) {
				repr += "^"
				represented = true
			}
			if !represented && m.beamY == y {
				if slices.Contains(m.beamX, x) {
					repr += "|"
					represented = true
				}
			}
			if !represented {
				repr += "."
			}
		}
		repr += "\n"
	}
	return repr
}

func NewTachyonManifold(width int, height int, splitterPositions *[]core.Pos, beamPosition *core.Pos) *TachyonManifold {
	size := max(width*height/64, 1)
	bitArray := core.NewBitArray(size)
	bitMap := core.NewBitMap(bitArray, width, height)
	for _, pos := range *splitterPositions {
		bitMap.Set(pos.X, pos.Y)
	}

	beamX := make([]int, 1, startingBeamCapacity)
	beamX[0] = beamPosition.X

	return &TachyonManifold{
		manifold: NewManifold(core.NewBitMap(bitArray, width, height)),
		beamY:    beamPosition.Y,
		beamX:    beamX,
		nSplits:  0,
	}
}

func (m *TachyonManifold) Step() error {
	m.beamY++
	if m.beamY == m.manifold.height {
		return ErrBeamLeftManifold
	}
	newBeams := make([]int, 0, startingBeamCapacity)
	for _, beamX := range m.beamX {
		if m.manifold.splitters.IsSet(beamX, m.beamY) {
			m.nSplits++

			// split beam
			if beamX-1 >= 0 && beamX-1 < m.manifold.width && !slices.Contains(newBeams, beamX-1) {
				newBeams = append(newBeams, beamX-1)
			}
			if beamX+1 >= 0 && beamX+1 < m.manifold.width && !slices.Contains(newBeams, beamX+1) {
				newBeams = append(newBeams, beamX+1)
			}
		} else {
			if !slices.Contains(newBeams, beamX) {
				newBeams = append(newBeams, beamX)
			}
		}
	}
	m.beamX = newBeams
	return nil
}

func ReadTachyonManifold(bfr *bufio.Reader) (*TachyonManifold, error) {
	currX := 0
	currY := 0
	width := -1
	height := -1
	var beamPosition *core.Pos = nil
	splitterPositions := make([]core.Pos, 0, startingSplitterCapacity)

	for {
		b, err := bfr.ReadByte()
		if err == io.EOF {
			height = currY
			return NewTachyonManifold(width, height, &splitterPositions, beamPosition), nil
		} else if err != nil {
			return nil, err
		}

		switch b {
		case '\n':
			if width == -1 {
				width = currX
			} else if width != currX {
				return nil, fmt.Errorf("width changed during reading from %d to %d", width, currX)
			}
			currX = 0
			currY++
			continue

		case 'S':
			if currY != 0 {
				return nil, fmt.Errorf("beam start position starting at unexpected point (%d, %d)", currX, currY)
			}
			if beamPosition != nil {
				return nil, errors.New("manifold has more than one beam start position")
			}
			beamPosition = core.NewPos(currX, currY)
		case '^':
			splitterPositions = append(splitterPositions, *core.NewPos(currX, currY))
		}
		currX++
	}
}

func PartOne() {
	file := core.ReadPackageData("day7", "input.dat")
	bfr := bufio.NewReader(file)
	manifold, err := ReadTachyonManifold(bfr)
	if err != nil {
		panic(err)
	}

	for {
		err := manifold.Step()
		if err == ErrBeamLeftManifold {
			break
		}
		core.Check(err)
	}

	fmt.Printf("Day 7 - Part One: %d\n", manifold.nSplits)
}
