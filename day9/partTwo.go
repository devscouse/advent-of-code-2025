package day9

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"slices"

	"github.com/devscouse/advent-of-code-2025/core"
)

var directions [4][2]int = [4][2]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

type Edge struct {
	xMin int
	xMax int
	yMin int
	yMax int
}

func minMax(a, b int) (int, int) {
	return min(a, b), max(a, b)
}

func CompressPositions(redPositions *[]core.Pos) *map[int]int {
	values := make([]int, 0, len(*redPositions))
	for _, p := range *redPositions {
		values = append(values, p.X)
		values = append(values, p.Y)
	}

	slices.Sort(values)
	encodeMap := make(map[int]int)
	decodeMap := make(map[int]int)
	for i, v := range values {
		_, ok := encodeMap[v]
		if !ok {
			encodeMap[v] = i
			decodeMap[i] = v
		}
	}

	for i, p := range *redPositions {
		p.X = encodeMap[p.X]
		p.Y = encodeMap[p.Y]
		(*redPositions)[i] = p
	}

	return &decodeMap
}

func GetGreenTileEdges(redPositions *[]core.Pos) *[]Edge {
	edges := make([]Edge, 0, 1000)
	for i, thisPos := range *redPositions {

		beforeIdx := i - 1
		if beforeIdx < 0 {
			beforeIdx = len(*redPositions) - 1
		}
		beforePos := (*redPositions)[beforeIdx]
		xMin, xMax := minMax(thisPos.X, beforePos.X)
		yMin, yMax := minMax(thisPos.Y, beforePos.Y)
		edge := Edge{xMin: xMin, xMax: xMax, yMin: yMin, yMax: yMax}
		edges = append(edges, edge)
	}
	return &edges
}

func PointInPolygon(x, y int, edges *[]Edge) bool {
	inside := false

	for _, e := range *edges {
		if e.xMin == e.xMax || y < e.yMin || y > e.yMax {
			continue
		}
		// Check if horizontal crosses edge or vertical crosses edge
		inside = !inside
	}
	return inside
}

func BoxInPolygon(a, b *core.Pos, edges *[]Edge) bool {
	xMin, xMax := minMax(a.X, b.X)
	yMin, yMax := minMax(a.Y, b.Y)

	// Check other corners
	corners := []core.Pos{
		{X: a.X, Y: b.Y},
		{X: b.X, Y: a.Y},
	}

	for _, corner := range corners {
		if !PointInPolygon(corner.X, corner.Y, edges) {
			return false
		}
	}

	// Check if any edge intersects the box
	for _, e := range *edges {
		if xMin < e.xMax && xMax > e.xMin && yMin < e.yMax && yMax > e.yMin {
			return false
		}
	}

	return true
}

func GetGreenTileEdgeBitMap(redPositions *[]core.Pos) *core.BitMap {
	width, height := 0, 0
	for _, p := range *redPositions {
		width = max(width, p.X+1)
		height = max(height, p.Y+1)
	}

	size := width * height
	bitMap := core.NewBitMap(core.NewBitArray(size), width, height)

	for i, thisPos := range *redPositions {

		beforeIdx := i - 1
		if beforeIdx < 0 {
			beforeIdx = len(*redPositions) - 1
		}
		beforePos := (*redPositions)[beforeIdx]
		xMin, xMax := minMax(thisPos.X, beforePos.X)
		yMin, yMax := minMax(thisPos.Y, beforePos.Y)

		for y := yMin; y <= yMax; y++ {
			for x := xMin; x <= xMax; x++ {
				bitMap.Set(x, y)
			}
		}
	}
	return bitMap
}

type Area struct {
	a    core.Pos
	b    core.Pos
	area int64
}

func AbsInt64(value int64) int64 {
	if value < 0 {
		return value * -1
	}
	return value
}

func CalculateArea(a, b *core.Pos) int64 {
	dx := AbsInt64(int64(a.X-b.X)) + 1
	dy := AbsInt64(int64(a.Y-b.Y)) + 1
	return dx * dy
}

func GetAllAreas(redPositions *[]core.Pos, decodeMap *map[int]int) *[]Area {
	n := len(*redPositions)
	areas := make([]Area, 0, n*n)
	for i, p1 := range *redPositions {
		for j := i + 1; j < len(*redPositions); j++ {
			p2 := (*redPositions)[j]
			areas = append(
				areas,
				Area{
					a: p1,
					b: p2,
					area: CalculateArea(
						core.NewPos((*decodeMap)[p1.X], (*decodeMap)[p1.Y]),
						core.NewPos((*decodeMap)[p2.X], (*decodeMap)[p2.Y]),
					),
				},
			)
		}
	}
	return &areas
}

func ReturnFirstValidArea(areas *[]Area, edges *[]Edge) Area {
	for _, area := range *areas {
		if BoxInPolygon(&area.a, &area.b, edges) {
			return area
		}
	}
	panic(errors.New("no Area is valid"))
}

func PartTwo() {
	file := core.ReadPackageData("day9", "input.dat")
	bfr := bufio.NewReader(file)
	positions := ReadPositions(bfr)
	log.Printf("%d Positions\n", len(*positions))

	decodeMap := CompressPositions(positions)
	log.Printf("%d positions compressed (%d encodings)", len(*positions), len(*decodeMap))

	edges := GetGreenTileEdges(positions)
	log.Printf("%d Edges\n", len(*edges))

	areas := GetAllAreas(positions, decodeMap)
	slices.SortFunc(*areas, func(a Area, b Area) int { return int(int64(b.area) - int64(a.area)) })

	area := ReturnFirstValidArea(areas, edges)

	log.Printf("Largest valid area %+v", area)
	fmt.Printf("Day 9 - Part Two: %d\n", area.area)
}
