/*Package day8*/
package day8

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"maps"
	"slices"
	"strconv"
	"strings"

	"github.com/devscouse/advent-of-code-2025/core"
)

const startJunctionCapacity = 30

type JunctionPair struct {
	idxOne   int
	idxTwo   int
	distance float64
}

func ReadJunctionPosition(bfr *bufio.Reader) (*core.Vector3, error) {
	line, _ := bfr.ReadString('\n')
	if len(line) == 0 {
		return nil, io.EOF
	}
	parts := strings.Split(line[:len(line)-1], ",")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid junction position: %s", line)
	}
	x, err := strconv.ParseFloat(parts[0], 64)
	core.Check(err)

	y, err := strconv.ParseFloat(parts[1], 64)
	core.Check(err)

	z, err := strconv.ParseFloat(parts[2], 64)
	core.Check(err)

	return core.NewVector3(x, y, z), nil
}

func ReadJunctionPositions(bfr *bufio.Reader) *[]core.Vector3 {
	junctionPositions := make([]core.Vector3, 0, startJunctionCapacity)
	for {
		position, err := ReadJunctionPosition(bfr)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		junctionPositions = append(junctionPositions, *position)
	}
	return &junctionPositions
}

func GetJunctionPairs(positions *[]core.Vector3) *[]JunctionPair {
	nJunctions := len(*positions)
	pairs := make([]JunctionPair, 0, nJunctions*nJunctions)

	for i, pos := range *positions {
		for j := i + 1; j < nJunctions; j++ {
			dist := pos.EuclideanDistance(&(*positions)[j])
			pair := JunctionPair{idxOne: i, idxTwo: j, distance: dist}
			pairs = append(pairs, pair)
		}
	}
	return &pairs
}

func SortPairsClosestFirst(pairs *[]JunctionPair) {
	slices.SortFunc(*pairs, func(a JunctionPair, b JunctionPair) int {
		if a.distance < b.distance {
			return -1
		} else {
			return 1
		}
	})
}

func PartOne() {
	file := core.ReadPackageData("day8", "input.dat")
	bfr := bufio.NewReader(file)
	positions := ReadJunctionPositions(bfr)
	log.Printf("%d positions loaded\n", len(*positions))

	pairs := GetJunctionPairs(positions)
	log.Printf("%d pairs created\n", len(*pairs))

	SortPairsClosestFirst(pairs)
	log.Printf("%d pairs sorted\n", len(*pairs))

	circuitPointers := make([]*core.Set, len(*positions))
	for i := range *positions {
		newSet := core.NewSet()
		newSet.Add(i)
		circuitPointers[i] = newSet
	}

	for i, pair := range *pairs {
		if i >= 1000 {
			break
		}
		circuitOne := circuitPointers[pair.idxOne]
		circuitTwo := circuitPointers[pair.idxTwo]

		if circuitOne == circuitTwo {
			continue
		}
		circuitOne = circuitOne.UnionInPlace(circuitTwo)
		for junctionIdx := range *circuitTwo {
			circuitPointers[junctionIdx] = circuitOne
		}
	}

	circuitSizes := make(map[*core.Set]int, len(circuitPointers))
	for _, circuit := range circuitPointers {
		circuitSizes[circuit] = len(*circuit)
	}
	circuitSizesArr := slices.Collect(maps.Values(circuitSizes))

	b, err := json.MarshalIndent(circuitSizesArr, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b) + "\n")

	slices.Sort(circuitSizesArr)
	slices.Reverse(circuitSizesArr)

	total := 1
	for i := range 3 {
		total *= circuitSizesArr[i]
	}

	fmt.Printf("Day 8 - Part One: %d\n", total)
}
