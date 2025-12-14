package day4

import (
	"bufio"
	"fmt"

	"github.com/devscouse/advent-of-code-2025/common"
)

// RemovePossibleRolls modifies the passed common.BitMap, unsetting any bits that have
// fewer than 4 surrounding set bits. The number of bits unset is returned.
func RemovePossibleRolls(b *common.BitMap) int {
	count := 0
	for x := range b.Width {
		for y := range b.Height {
			if b.IsSet(x, y) && b.CountSurroundingSet(x, y) < 4 {
				b.Unset(x, y)
				count++
			}
		}
	}
	return count
}

func PartTwo() {
	file := common.ReadPackageData("day4", "input.dat")
	bfr := bufio.NewReader(file)
	bitMap := ReadMap(bfr)
	totalCount := 0
	count := 0
	for {
		count = RemovePossibleRolls(bitMap)
		if count == 0 {
			break
		}
		totalCount += count
	}
	fmt.Printf("Day 4 - Part Two: %d\n", totalCount)
}
