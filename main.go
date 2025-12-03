package main

import (
	"log"
	"os"
	"strconv"

	"github.com/devscouse/advent-of-code-2025/common"
	"github.com/devscouse/advent-of-code-2025/day1"
	"github.com/devscouse/advent-of-code-2025/day2"
)

type (
	SolverFunc func()
	Solvers    map[int]map[int]SolverFunc
)

var solvers Solvers = make(Solvers)

func (s Solvers) addSolver(dayNumber int, partNumber int, solver SolverFunc) {
	if solvers[dayNumber] == nil {
		solvers[dayNumber] = make(map[int]SolverFunc)
	}
	solvers[dayNumber][partNumber] = solver
}

func (s Solvers) runSolver(dayNumber int, partNumber int) {
	if solvers[dayNumber] == nil || solvers[dayNumber][partNumber] == nil {
		log.Printf("No Solver is available for day %d part %d\n", dayNumber, partNumber)
		os.Exit(1)
	}
	solvers[dayNumber][partNumber]()
}

func init() {
	solvers.addSolver(1, 1, day1.SolvePartOne)
	solvers.addSolver(1, 2, day1.SolvePartTwo)
	solvers.addSolver(2, 1, day2.PartOne)
	solvers.addSolver(2, 2, day2.PartTwo)
}

func main() {
	if len(os.Args) < 3 {
		log.Println("Expected 2 command line arguments. The day number and the part number")
		os.Exit(1)
	}

	dayNumber, err := strconv.Atoi(os.Args[1])
	common.Check(err)

	partNumber, err := strconv.Atoi(os.Args[2])
	common.Check(err)

	solvers.runSolver(dayNumber, partNumber)

	solver := solvers[dayNumber][partNumber]
	solver()
}
