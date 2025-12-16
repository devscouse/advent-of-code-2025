package day8

import (
	"bufio"
	"fmt"
	"log"

	"github.com/devscouse/advent-of-code-2025/common"
)

func PartTwo() {
	file := common.ReadPackageData("day8", "input.dat")
	bfr := bufio.NewReader(file)
	positions := ReadJunctionPositions(bfr)
	log.Printf("%d positions loaded\n", len(*positions))

	pairs := GetJunctionPairs(positions)
	log.Printf("%d pairs created\n", len(*pairs))

	SortPairsClosestFirst(pairs)
	log.Printf("%d pairs sorted\n", len(*pairs))

	circuitPointers := make([]*common.Set, len(*positions))
	for i := range *positions {
		newSet := common.NewSet()
		newSet.Add(i)
		circuitPointers[i] = newSet
	}

	nJunctions := len(*positions)

	for _, pair := range *pairs {
		circuitOne := circuitPointers[pair.idxOne]
		circuitTwo := circuitPointers[pair.idxTwo]

		if circuitOne == circuitTwo {
			continue
		}
		circuitOne = circuitOne.UnionInPlace(circuitTwo)
		if len(*circuitOne) == nJunctions {
			posOne := (*positions)[pair.idxOne]
			posTwo := (*positions)[pair.idxTwo]
			fmt.Printf("Day 8 - Part Two: %.0f\n", posOne.X*posTwo.X)
			break
		}
		for junctionIdx := range *circuitTwo {
			circuitPointers[junctionIdx] = circuitOne
		}
	}
}
