/*Package day3*/
package day3

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/devscouse/advent-of-code-2025/common"
)

func getMaxJoltageFrom2Batteries(bank string) int {
	maxMajorJolt := '0'
	minMajorJolt := '0'

	for i, r := range bank {
		if i < len(bank)-1 && r > maxMajorJolt {
			maxMajorJolt = r
			minMajorJolt = '0'
		} else if r > minMajorJolt {
			minMajorJolt = r
		}
	}

	major, err := strconv.Atoi(string(maxMajorJolt))
	common.Check(err)

	minor, err := strconv.Atoi(string(minMajorJolt))
	common.Check(err)

	maxJoltage := major*10 + minor
	fmt.Printf("Max Joltage: %d\n", maxJoltage)
	return maxJoltage
}

func LoadNextBatteryBank(bfr *bufio.Reader) (string, error) {
	str, err := bfr.ReadString('\n')
	if err != nil {
		return "", err
	}
	return str[:len(str)-1], nil
}

func PartOne() {
	path := filepath.Join(".", "day3", "data", "input.dat")
	file, err := os.Open(path)
	common.Check(err)

	bfr := bufio.NewReader(file)
	totalJoltage := 0

	for {
		bank, err := LoadNextBatteryBank(bfr)
		if bank == "" {
			break
		}
		if err != nil && err != io.EOF {
			log.Fatalf("%e\n", err)
			break
		}
		totalJoltage += getMaxJoltageFrom2Batteries(bank)

		if err == io.EOF {
			break
		}
	}

	fmt.Printf("Day 3 - Part One: %d\n", totalJoltage)
}
