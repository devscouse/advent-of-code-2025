package day6

import (
	"bufio"
	"fmt"
	"io"
	"log"

	"github.com/devscouse/advent-of-code-2025/common"
)

func readOperandSection(bfr *bufio.Reader) *[]string {
	operandSection := make([]string, 0, 3)
	for {
		bytes, err := bfr.Peek(1)
		common.Check(err)
		if bytes[0] == '*' || bytes[0] == '+' {
			return &operandSection
		}

		s, err := bfr.ReadString('\n')
		common.Check(err)
		operandSection = append(operandSection, s[:len(s)-1])
	}
}

func ReadCephalopodMathProblems(bfr *bufio.Reader) *[]MathProblem {
	operandSection := readOperandSection(bfr)

	maxLen := 0
	for _, line := range *operandSection {
		maxLen = max(maxLen, len(line))
	}

	problems := make([]MathProblem, 0, 25)
	problem := NewMathProblem(3, '?')

	operandIdx := 0

	for column := range maxLen {
		log.Printf("Processing column %d...\n", column)
		dividerFound := true
		operand := 0

		for _, line := range *operandSection {

			if column > len(line) {
				dividerFound = false
				continue
			}

			b := line[column]

			if '0' <= b && b <= '9' {
				operand = common.AddLeastSignificantDigit(operand, int(b-'0'))
				dividerFound = false
			}
		}

		if dividerFound && len(problem.operands) > 0 {
			log.Printf("problem operand parsing completed: %v\n", problem)
			problems = append(problems, *problem)
			problem = NewMathProblem(3, '?')
			operandIdx = 0
		} else if operand > 0 {
			problem.AddOperand(operand)
			operandIdx++
		}
	}
	if len(problem.operands) > 0 {
		log.Printf("problem operand parsing completed: %v\n", problem)
		problems = append(problems, *problem)
	}

	// Read operators
	problemIdx := 0
	for {
		r, _, err := bfr.ReadRune()
		if err == io.EOF {
			break
		}
		common.Check(err)

		if r == ' ' || r == '\n' {
			continue
		}

		if r == '*' || r == '+' {
			problems[problemIdx].operator = byte(r)
			problemIdx++
		}
	}

	return &problems
}

func PartTwo() {
	file := common.ReadPackageData("day6", "input.dat")
	bfr := bufio.NewReader(file)
	mathProblems := ReadCephalopodMathProblems(bfr)
	fmt.Printf("mathProblems: %+v\n", mathProblems)

	solution := 0
	for _, p := range *mathProblems {
		solution += p.Solve()
	}
	fmt.Printf("Day 6 - Part Two: %d\n", solution)
}
