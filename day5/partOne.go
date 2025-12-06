/*Package day5*/
package day5

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/devscouse/advent-of-code-2025/common"
)

type IDRange struct {
	min uint64
	max uint64
}

func (r *IDRange) InRange(id uint64) bool {
	return r.min <= id && id <= r.max
}

func (r *IDRange) Overlaps(other *IDRange) bool {
	return ((other.min <= r.min && r.min <= other.max) || (other.min <= r.max && r.max <= other.max) || (r.min <= other.min && other.min <= r.max) || (r.min <= other.max && other.max <= r.max))
}

func (r *IDRange) Merge(other *IDRange) *IDRange {
	r.min = min(r.min, other.min)
	r.max = max(r.max, other.max)
	return r
}

const (
	queryIngredientIDsStartingSize      = 1000
	freshIngredientIDRangesStartingSize = 100
)

func ReadFreshIngredientIdsRange(bfr *bufio.Reader, freshIngredients *[]IDRange) error {
	// Peek ahead to see if we have reached the blank line
	peek, err := bfr.Peek(1)
	common.Check(err)

	// If we have, consume the blank line then return an EOF
	if peek[0] == '\n' {
		_, err := bfr.ReadByte()
		common.Check(err)
		return io.EOF
	}

	str, err := bfr.ReadString('-')
	common.Check(err)
	str = str[:len(str)-1]

	minID, err := strconv.ParseUint(str, 10, 64)
	common.Check(err)

	str, err = bfr.ReadString('\n')
	common.Check(err)
	str = str[:len(str)-1]

	maxID, err := strconv.ParseUint(str, 10, 64)
	common.Check(err)

	idRange := IDRange{min: minID, max: maxID}
	for i, r := range *freshIngredients {
		if r.Overlaps(&idRange) {
			log.Printf("Merging ranges: %+v and %+v\n", r, idRange)
			(*freshIngredients)[i] = *r.Merge(&idRange)
			break
		}
	}

	*freshIngredients = append(*freshIngredients, idRange)
	return nil
}

func ReadFreshIngredientIds(bfr *bufio.Reader) *[]IDRange {
	freshIngredients := make([]IDRange, 0, freshIngredientIDRangesStartingSize)
	rowNumber := 0
	for {
		rowNumber++
		err := ReadFreshIngredientIdsRange(bfr, &freshIngredients)
		log.Printf("Database row %d read\n", rowNumber)
		if err != nil {
			break
		}
	}
	return &freshIngredients
}

func ReadQueryIngredientIds(bfr *bufio.Reader) []uint64 {
	ingredientIDs := make([]uint64, 0, queryIngredientIDsStartingSize)
	endReached := false
	for !endReached {
		str, err := bfr.ReadString('\n')
		if err != nil && err != io.EOF {
			continue
		} else if err == nil {
			str = str[:len(str)-1]
		} else {
			endReached = true
		}

		id, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			log.Printf("Invalid ID: '%s'\n", str)
			continue
		}
		ingredientIDs = append(ingredientIDs, id)
	}
	return ingredientIDs
}

func CountFreshIngredients(query []uint64, fresh []IDRange) int {
	count := 0
	for _, q := range query {
		for _, r := range fresh {
			if r.min <= q && q <= r.max {
				count++
				break
			}
		}
	}
	return count
}

func PartOne() {
	file := common.ReadPackageData("day5", "input.dat")
	bfr := bufio.NewReader(file)
	log.Println("File opened")
	freshIngredients := ReadFreshIngredientIds(bfr)
	fmt.Printf("%v\n", freshIngredients)
	queryIngredients := ReadQueryIngredientIds(bfr)
	fmt.Printf("%v\n", queryIngredients)

	count := CountFreshIngredients(queryIngredients, *freshIngredients)
	fmt.Printf("Day 5 - Part One: %d\n", count)
}
