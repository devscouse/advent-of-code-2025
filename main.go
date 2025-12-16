package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/devscouse/advent-of-code-2025/common"
	"github.com/devscouse/advent-of-code-2025/day1"
	"github.com/devscouse/advent-of-code-2025/day2"
	"github.com/devscouse/advent-of-code-2025/day3"
	"github.com/devscouse/advent-of-code-2025/day4"
	"github.com/devscouse/advent-of-code-2025/day5"
	"github.com/devscouse/advent-of-code-2025/day6"
	"github.com/devscouse/advent-of-code-2025/day7"
	"github.com/devscouse/advent-of-code-2025/day8"
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

	start := time.Now()
	solvers[dayNumber][partNumber]()
	runtime := time.Since(start).Seconds()
	fmt.Printf("Runtime: %.2f seconds (%.2f ms)\n", runtime, runtime*1000)
}

func init() {
	solvers.addSolver(1, 1, day1.SolvePartOne)
	solvers.addSolver(1, 2, day1.SolvePartTwo)
	solvers.addSolver(2, 1, day2.PartOne)
	solvers.addSolver(2, 2, day2.PartTwo)
	solvers.addSolver(3, 1, day3.PartOne)
	solvers.addSolver(3, 2, day3.PartTwo)
	solvers.addSolver(4, 1, day4.PartOne)
	solvers.addSolver(4, 2, day4.PartTwo)
	solvers.addSolver(5, 1, day5.PartOne)
	solvers.addSolver(5, 2, day5.PartTwo)
	solvers.addSolver(6, 1, day6.PartOne)
	solvers.addSolver(6, 2, day6.PartTwo)
	solvers.addSolver(7, 1, day7.PartOne)
	solvers.addSolver(7, 2, day7.PartTwo)
	solvers.addSolver(8, 1, day8.PartOne)
	solvers.addSolver(8, 2, day8.PartTwo)
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
}
