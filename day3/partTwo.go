package day3

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"

	"github.com/devscouse/advent-of-code-2025/core"
)

func getTotalJoltage(enabledBatteries []rune) uint64 {
	totalJoltage := uint64(0)
	factor := 1
	slices.Reverse(enabledBatteries)
	for _, j := range enabledBatteries {
		joltage, err := strconv.Atoi(string(j))
		core.Check(err)

		totalJoltage += uint64(joltage * factor)
		factor *= 10
	}
	return totalJoltage
}

func getMaxJoltage(bank string, maxEnabled int) uint64 {
	log.Printf("Battery Bank: %s\n", bank)
	enabledBatteries := make([]rune, maxEnabled)
	for i := range enabledBatteries {
		enabledBatteries[i] = '0'
	}
	nBatteries := len(bank)

	for i, this := range bank {
		// log.Printf("Checking Battery %3d (Joltage: %c) | %s\n", i, this, string(enabledBatteries))
		enabled := false
		batteriesLeft := nBatteries - i
		for j, curr := range enabledBatteries {

			// If we have enabled battery i we need to clear every other
			// enabled battery after the position it was stored into
			if enabled {
				enabledBatteries[j] = '0'
				continue
			}

			// Check that we can use battery i at position j and still enabled
			// the remaining batteries we need to
			// E.g. if we can enable 5 batteries out of 10, we cannot enable battery 8 at
			// position 3
			batteriesToEnable := maxEnabled - j
			if batteriesToEnable > batteriesLeft {
				continue
			}

			if this > curr {
				enabledBatteries[j] = this
				enabled = true
				// log.Printf("Enabling Battery %3d at position %d\n", i, j)
			}
		}
	}

	log.Printf("Enabled batteries: %s\n", string(enabledBatteries))
	joltage := getTotalJoltage(enabledBatteries)
	return joltage
}

func PartTwo() {
	path := filepath.Join(".", "day3", "data", "input.dat")
	file, err := os.Open(path)
	core.Check(err)

	bfr := bufio.NewReader(file)
	totalJoltage := uint64(0)

	for {
		bank, err := LoadNextBatteryBank(bfr)
		if bank == "" {
			break
		}
		if err != nil && err != io.EOF {
			log.Fatalf("%e\n", err)
			break
		}
		totalJoltage += getMaxJoltage(bank, 12)

		if err == io.EOF {
			break
		}
	}

	fmt.Printf("Day 3 - Part Two: %d\n", totalJoltage)
}
