package day2

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"

	"github.com/devscouse/advent-of-code-2025/core"
)

func Pow10(n uint) uint {
	return uint(math.Pow(10, float64(n)))
}

func repeatSplit(split uint, n uint) uint {
	factor := Pow10(countDigits(split))
	result := split
	for range n {
		result = result*factor + split
	}
	return result
}

func takeLeftDigits(value uint, n uint) uint {
	nDigits := countDigits(value)
	if nDigits == n {
		return value
	}
	divisor := Pow10(nDigits - n)
	return value / divisor
}

func sumPartTwoInvalidIDsInRange(from uint, to uint) uint {
	var sum uint = 0

	maxDigits := countDigits(to)
	minDigits := countDigits(from)

	// 1. Take left digits
	// 2. Repeat left digits up until ndigits
	// 3. If value is in range, add to sum
	// 4. Increment the left digits
	invalidIds := make(map[uint]bool)
	for nDigits := minDigits; nDigits <= maxDigits; nDigits++ {
		checkedDigitCounts := []uint{}

		for leftDigitCount := nDigits / 2; leftDigitCount > 0; leftDigitCount-- {
			// log.Printf("Looking for invalid IDs with %d digits a repeating pattern of %d numbers", nDigits, leftDigitCount)
			if nDigits%leftDigitCount != 0 {
				// log.Printf("Pattern cannot repeat to reach %d digits", nDigits)
				continue
			}

			skip := false
			for _, checkedDigitCount := range checkedDigitCounts {
				if checkedDigitCount%leftDigitCount == 0 {
					// log.Printf("We have checked %d already (a multiple of %d)\n", checkedDigitCount, leftDigitCount)
					skip = true
					break
				}
			}
			if skip {
				continue
			}

			repeatAmount := (nDigits - leftDigitCount) / leftDigitCount

			leftDigits := Pow10(leftDigitCount - 1)
			end := Pow10(leftDigitCount)

			for leftDigits < end {

				id := repeatSplit(leftDigits, repeatAmount)

				if id > to {
					// log.Printf("ID outside of range: %d\n", id)
					break
				}
				if from <= id && id <= to {
					if !invalidIds[id] {
						log.Printf("Invalid ID Found: %d\n", id)
						sum += id
						invalidIds[id] = true
					}
				}
				leftDigits++

			}
			checkedDigitCounts = append(checkedDigitCounts, leftDigitCount)
			log.Printf("checkedDigitCounts: %v\n", checkedDigitCounts)
		}
	}
	return sum
}

func PartTwo() {
	path := filepath.Join(".", "day2", "data", "input.dat")
	file, err := os.Open(path)
	core.Check(err)
	bfr := bufio.NewReader(file)
	var sum uint = 0
	for {
		left, right, err := ReadNextProductIDRange(bfr)
		if err != nil {
			fmt.Printf("%s\n", err)
			break
		}
		fmt.Printf("\nProduct ID Range: %d - %d\n", left, right)
		sum += sumPartTwoInvalidIDsInRange(left, right)
	}

	fmt.Printf("Day Two - Part Two: %d\n", sum)
	// 30962646867 is incorrect
}
