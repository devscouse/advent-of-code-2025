package main

import (
	"log"
	"os"
	"strconv"

	"github.com/devscouse/advent-of-code-2025/day1"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	if len(os.Args) < 3 {
		log.Println("Expected 2 command line arguments. The day number and the part number")
		os.Exit(1)
	}

	dayNumber, err := strconv.Atoi(os.Args[1])
	check(err)

	partNumber, err := strconv.Atoi(os.Args[2])
	check(err)

	if dayNumber == 1 && partNumber == 1 {
		day1.SolvePartOne()
	} else if dayNumber == 1 && partNumber == 2 {
		day1.SolvePartTwo()
	}
}
