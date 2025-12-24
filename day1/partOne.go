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

const (
	MaxDialValue     = 99
	InitialDialValue = 50
)

func turnDialLeft(start int, amount int) int {
	newValue := start - amount
	if newValue >= 0 {
		return newValue
	}
	return turnDialLeft(99, (newValue*-1)-1)
}

func turnDialRight(start int, amount int) int {
	newValue := start + amount
	if newValue <= MaxDialValue {
		return newValue
	}
	return turnDialRight(0, newValue-MaxDialValue-1)
}

func solve(instructions string) int {
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

		switch direction {
		case 'L':
			dialValue = turnDialLeft(dialValue, amount)
		case 'R':
			dialValue = turnDialRight(dialValue, amount)
		default:
			os.Exit(1)
		}

		log.Printf("Instruction %d: dialValue=%d\n", i, dialValue)

		if dialValue == 0 {
			zeroCounts++
		}
	}

	return zeroCounts
}

func SolvePartOne() {
	path := filepath.Join(".", "day1", "data", "input.dat")
	data, err := os.ReadFile(path)
	core.Check(err)
	solution := solve(string(data))
	fmt.Printf("Day 1 - Part One Solution: %d\n", solution)
}
