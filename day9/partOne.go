/*Package day9*/
package day9

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"

	"github.com/devscouse/advent-of-code-2025/core"
)

const startPositionsCapacity = 1000

func ReadPosition(bfr *bufio.Reader) (*core.Pos, error) {
	str, err := bfr.ReadString('\n')
	if (err != nil && err != io.EOF) || len(str) <= 1 {
		return nil, err
	}
	if str[len(str)-1] == '\n' {
		str = str[:len(str)-1]
	}

	parts := strings.Split(str, ",")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid position: %s", str)
	}

	x, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}

	y, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	return core.NewPos(x, y), nil
}

func ReadPositions(bfr *bufio.Reader) *[]core.Pos {
	positions := make([]core.Pos, 0, startPositionsCapacity)
	for {
		pos, err := ReadPosition(bfr)
		if err != nil {
			return &positions
		}
		positions = append(positions, *pos)
	}
}

func GetArea(a, b *core.Pos) float64 {
	return (math.Abs(float64(a.X-b.X)) + 1) * (math.Abs(float64(a.Y-b.Y)) + 1)
}

func FindLargestArea(positions *[]core.Pos) float64 {
	largestArea := 0.0
	for i, posOne := range *positions {
		for j := i + 1; j < len(*positions); j++ {
			posTwo := (*positions)[j]
			largestArea = max(largestArea, GetArea(&posOne, &posTwo))
		}
	}
	return largestArea
}

func PartOne() {
	file := core.ReadPackageData("day9", "input.dat")
	bfr := bufio.NewReader(file)
	positions := ReadPositions(bfr)

	fmt.Printf("Day 9 - Part One: %.0f\n", FindLargestArea(positions))
}
