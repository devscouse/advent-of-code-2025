// Package day6
package day6

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/devscouse/advent-of-code-2025/common"
)

const defaultMathProblemsCapacity = 100

type MathProblem struct {
	operands []int
	operator byte
}

func NewMathProblem(operandSize int, operator byte) *MathProblem {
	operands := make([]int, 0, operandSize)
	return &MathProblem{operands: operands, operator: operator}
}

func (p *MathProblem) AddOperand(operand int) {
	p.operands = append(p.operands, operand)
}

func (p *MathProblem) Solve() int {
	if len(p.operands) == 0 {
		panic(errors.New("cannot solve MathProblem with no operands"))
	}

	solution := p.operands[0]
	for i := 1; i < len(p.operands); i++ {
		switch p.operator {
		case '+':
			solution += p.operands[i]
		case '*':
			solution *= p.operands[i]
		default:
			panic(fmt.Errorf("operand %c is not supported", p.operator))
		}
	}
	return solution
}

func ReadOperands(bfr *bufio.Reader, capacity int) []int {
	operands := make([]int, 0, capacity)
	currOperand := ""
	for {
		r, _, err := bfr.ReadRune()
		if err == io.EOF {
			break
		}

		if '0' <= r && r <= '9' {
			currOperand += string(r)
			continue
		}

		if r == '\n' || r == ' ' {
			if len(currOperand) > 0 {
				value, err := strconv.Atoi(currOperand)
				common.Check(err)
				operands = append(operands, value)
				currOperand = ""
			}
		} else {
			panic(fmt.Errorf("unexpected character %c", r))
		}

		if r == '\n' {
			break
		}
	}
	return operands
}

func ReadOperators(bfr *bufio.Reader, problems *[]MathProblem) {
	currIdx := 0
	for {
		b, err := bfr.ReadByte()
		if err == io.EOF {
			break
		}

		if b == '*' || b == '+' {
			(*problems)[currIdx].operator = b
			currIdx++
			continue
		}

		if b == '\n' {
			return
		}
	}
}

func ReadMathProblems(bfr *bufio.Reader) *[]MathProblem {
	operands := ReadOperands(bfr, defaultMathProblemsCapacity)
	nProblems := len(operands)
	mathProblems := make([]MathProblem, nProblems)
	for i, operand := range operands {
		mathProblems[i] = MathProblem{operands: make([]int, 0), operator: '-'}
		mathProblems[i].AddOperand(operand)
	}

	for {
		bytes, err := bfr.Peek(1)
		if err == io.EOF {
			break
		}
		common.Check(err)
		if bytes[0] == '*' || bytes[0] == '+' {
			ReadOperators(bfr, &mathProblems)
		} else {
			operands = ReadOperands(bfr, nProblems)
			for i, operand := range operands {
				mathProblems[i].AddOperand(operand)
			}
		}
	}

	return &mathProblems
}

func PartOne() {
	file := common.ReadPackageData("day6", "input.dat")
	bfr := bufio.NewReader(file)
	mathProblems := ReadMathProblems(bfr)
	fmt.Printf("mathProblems: %+v\n", mathProblems)

	solution := 0
	for _, p := range *mathProblems {
		solution += p.Solve()
	}
	fmt.Printf("Day 6 - Part One: %d\n", solution)
}
