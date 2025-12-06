package day5

import (
	"bufio"
	"fmt"
	"log"

	"github.com/devscouse/advent-of-code-2025/common"
)

func CountAllFreshIngredients(freshIngredients *[]IDRange) int {
	count := 0
	for _, r := range *freshIngredients {
		count += int(r.max-r.min) + 1
	}
	return count
}

func MergeAllOverlappingRanges(ranges *[]IDRange) (*[]IDRange, int) {
	merges := 0
	newRanges := make([]IDRange, 0, freshIngredientIDRangesStartingSize)
	for _, r := range *ranges {
		merged := false
		for i, n := range newRanges {
			if n.Overlaps(&r) {
				newRanges[i] = *n.Merge(&r)
				merges++
				merged = true
				break
			}
		}

		if !merged {
			newRanges = append(newRanges, r)
		}
	}
	log.Printf("%d ranges after merging\n", len(newRanges))
	return &newRanges, merges
}

func PartTwo() {
	file := common.ReadPackageData("day5", "input.dat")
	bfr := bufio.NewReader(file)
	log.Println("File opened")
	freshIngredients := ReadFreshIngredientIds(bfr)
	log.Printf("%v\n", freshIngredients)

	merges := -1
	for merges != 0 {
		freshIngredients, merges = MergeAllOverlappingRanges(freshIngredients)
	}

	fmt.Printf("Day 5 - Part Two: %d\n", CountAllFreshIngredients(freshIngredients))
}
