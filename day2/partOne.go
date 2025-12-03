/*Package day2*/
package day2

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"github.com/devscouse/advent-of-code-2025/common"
)

func countDigits(value uint) uint {
	var ndigits uint = 1
	var divisor uint = 10
	for {
		if value/divisor == 0 {
			return ndigits
		}
		ndigits++
		divisor *= 10
	}
}

func SplitInt(value uint, left uint) (uint, uint) {
	divisor := uint(math.Pow(10, float64(left)))
	return value / divisor, value % divisor
}

func sumPartOneInvalidIDsInRange(from uint, to uint) uint {
	var sum uint = 0
	curr := from
	for curr <= to {
		nDigits := countDigits(curr)

		if nDigits%2 != 0 {
			curr = uint(math.Pow(10, float64(nDigits)))
			continue
		}

		splitAt := countDigits(curr) / 2
		leftSplit, _ := SplitInt(curr, splitAt)
		factor := uint(math.Pow(10, float64(splitAt)))
		invalidID := leftSplit*factor + leftSplit

		if invalidID >= from && invalidID <= to {
			log.Printf("Invalid ID (%d) found\n", invalidID)
			sum += invalidID
		}

		leftSplit++
		curr = leftSplit*factor + leftSplit
	}
	return sum
}

func ReadNextProductIDRange(bfr *bufio.Reader) (uint, uint, error) {
	leftStr, err := bfr.ReadString('-')
	if err != nil {
		return 0, 0, err
	}
	if leftStr[0] == '\n' {
		leftStr = leftStr[1:]
	}

	leftVal, err := strconv.Atoi(leftStr[:len(leftStr)-1])
	if err != nil {
		return 0, 0, err
	}

	rightStr, err := bfr.ReadString(',')
	if err != nil && err != io.EOF {
		return 0, 0, err
	}

	rightVal, err := strconv.Atoi(rightStr[:len(rightStr)-1])
	if err != nil {
		return 0, 0, err
	}

	return uint(leftVal), uint(rightVal), nil
}

func PartOne() {
	path := filepath.Join(".", "day2", "data", "input.dat")
	file, err := os.Open(path)
	common.Check(err)
	bfr := bufio.NewReader(file)
	var sum uint = 0
	for {
		left, right, err := ReadNextProductIDRange(bfr)
		if err != nil {
			fmt.Printf("%s\n", err)
			break
		}
		fmt.Printf("\nProduct ID Range: %d - %d\n", left, right)
		sum += sumPartOneInvalidIDsInRange(left, right)
	}

	fmt.Printf("Day Two - Part One: %d\n", sum)
}
