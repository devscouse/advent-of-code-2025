package day1

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/devscouse/advent-of-code-2025/core"
)

func turnDialLeftWithZeroCount(start int, amount int) (int, int) {
	newValue := start - amount
	if newValue > 0 {
		return newValue, 0
	} else if newValue == 0 {
		return newValue, 1
	}

	newValue, zeroCounts := turnDialLeftWithZeroCount(99, (newValue*-1)-1)
	if start == 0 {
		return newValue, zeroCounts
	}
	return newValue, zeroCounts + 1
}

func turnDialRightWithZeroCount(start int, amount int) (int, int) {
	newValue := start + amount
	if newValue <= MaxDialValue {
		return newValue, 0
	}
	newValue, zeroCounts := turnDialRightWithZeroCount(0, newValue-MaxDialValue-1)
	return newValue, zeroCounts + 1
}

func solvePartTwo(instructions string) int {
	dialValue := InitialDialValue
	zeroCounts := 0
	for i, instruction := range strings.Split(instructions, "\n") {
		if instruction == "" {
			continue
		}

		direction := instruction[0]
		amount, err := strconv.Atoi(instruction[1:])
		core.Check(err)

		log.Printf("Instruction %d: dialValue=%d direction=%c amount=%d\n", i, dialValue, direction, amount)
		turnZeroCounts := 0

		switch direction {
		case 'L':
			dialValue, turnZeroCounts = turnDialLeftWithZeroCount(dialValue, amount)
		case 'R':
			dialValue, turnZeroCounts = turnDialRightWithZeroCount(dialValue, amount)
		default:
			os.Exit(1)
		}

		log.Printf("Instruction %d: dialValue=%d, turnZeroCounts=%d\n", i, dialValue, turnZeroCounts)
		zeroCounts += turnZeroCounts
	}

	return zeroCounts
}

func SolvePartTwo() {
	path := filepath.Join(".", "day1", "data", "input.dat")
	data, err := os.ReadFile(path)
	core.Check(err)
	solution := solvePartTwo(string(data))
	fmt.Printf("Day 1 - Part Two Solution: %d\n", solution)
}
