package day7

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"

	"github.com/devscouse/advent-of-code-2025/core"
)

const startParticleCapacity = 1000

type QuantumTachyonManifold struct {
	manifold   *Manifold
	particlesY int
	particlesX map[int]int
}

func (q *QuantumTachyonManifold) String() string {
	repr := ""
	for y := range q.manifold.height {
		for x := range q.manifold.width {
			represented := false
			if q.manifold.splitters.IsSet(x, y) {
				repr += "^"
				represented = true
			}
			if !represented && q.particlesY == y {
				if q.particlesX[x] != 0 {
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

func NewQuantumTachyonManifold(width int, height int, splitterPositions *[]core.Pos, particleStart *core.Pos) *QuantumTachyonManifold {
	log.Printf("Creating new QuantumTachyonManifold with width=%d, height=%d, particleStart=%+v\n", width, height, particleStart)
	bitArray := core.NewBitArray(width * height)
	bitMap := core.NewBitMap(bitArray, width, height)
	for _, pos := range *splitterPositions {
		bitMap.Set(pos.X, pos.Y)
	}

	particlesX := make(map[int]int)
	particlesX[particleStart.X] = 1

	return &QuantumTachyonManifold{
		manifold:   NewManifold(core.NewBitMap(bitArray, width, height)),
		particlesY: particleStart.Y,
		particlesX: particlesX,
	}
}

func (q *QuantumTachyonManifold) Advance() error {
	if q.particlesY >= q.manifold.height {
		return ErrBeamLeftManifold
	}

	for particleX := range q.manifold.width {
		nTimelines := q.particlesX[particleX]
		if nTimelines == 0 {
			continue
		}
		if !q.manifold.splitters.IsSet(particleX, q.particlesY) {
			continue
		}
		q.particlesX[particleX] = 0
		if particleX-1 >= 0 {
			q.particlesX[particleX-1] += nTimelines
		}

		if particleX+1 < q.manifold.width {
			q.particlesX[particleX+1] += nTimelines
		}
	}
	q.particlesY++
	return nil
}

func ReadQuantumTachyonManifold(bfr *bufio.Reader) (*QuantumTachyonManifold, error) {
	currX := 0
	currY := 0
	width := -1
	height := -1
	var particleStart *core.Pos = nil
	splitterPositions := make([]core.Pos, 0, startingSplitterCapacity)

	for {
		b, err := bfr.ReadByte()
		if err == io.EOF {
			height = currY
			return NewQuantumTachyonManifold(width, height, &splitterPositions, particleStart), nil
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
				return nil, fmt.Errorf("particle start position starting at unexpected point (%d, %d)", currX, currY)
			}
			if particleStart != nil {
				return nil, errors.New("manifold has more than one particle start position")
			}
			particleStart = core.NewPos(currX, currY)
		case '^':
			splitterPositions = append(splitterPositions, *core.NewPos(currX, currY))
		}
		currX++
	}
}

func PartTwo() {
	file := core.ReadPackageData("day7", "input.dat")
	bfr := bufio.NewReader(file)
	manifold, err := ReadQuantumTachyonManifold(bfr)
	if err != nil {
		panic(err)
	}

	for {
		err := manifold.Advance()
		if err == ErrBeamLeftManifold {
			break
		}
		core.Check(err)
	}
	nTimelines := 0
	for _, positionTimelines := range manifold.particlesX {
		nTimelines += positionTimelines
	}
	fmt.Printf("Day 7 - Part Two: %d\n", nTimelines)
}
