/*Package day4*/
package day4

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/devscouse/advent-of-code-2025/common"
)

func ReadMap(bfr *bufio.Reader) *common.BitMap {
	currPos := 0
	mapWidth := 0
	mapHeight := 0
	bitArray := common.NewBitArray(64 * 500)

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
	return common.NewBitMap(bitArray, mapWidth, mapHeight)
}

func PartOne() {
	file := common.ReadPackageData("day4", "input.dat")
	bfr := bufio.NewReader(file)
	bitMap := ReadMap(bfr)
	count := 0

	for y := range bitMap.Height {
		for x := range bitMap.Width {
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
