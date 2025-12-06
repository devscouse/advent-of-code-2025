package day4

import (
	"bufio"
	"fmt"

	"github.com/devscouse/advent-of-code-2025/common"
)

func (b *BitArray) Unset(idx int) {
	word := idx / 64
	wordIdx := idx % 64
	if word >= len(b.bits) {
		return
	}
	b.bits[word] &^= (1 << wordIdx)
}

func (b *BitMap) GetPositionIdx(x int, y int) int {
	return y*b.width + x
}

func (b *BitMap) Unset(x int, y int) {
	b.array.Unset(b.GetPositionIdx(x, y))
}

// RemovePossibleRolls modifies the passed BitMap, unsetting any bits that have
// fewer than 4 surrounding set bits. The number of bits unset is returned.
func RemovePossibleRolls(b *BitMap) int {
	count := 0
	for x := range b.width {
		for y := range b.height {
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
