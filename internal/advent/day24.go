package advent

import (
	. "github.com/davidparks11/advent2021/internal/advent/day24"
)

type arithmeticLogicUnit struct {
	dailyProblem
}

func (a *arithmeticLogicUnit) Solve() interface{} {
	input := a.GetInputLines()
	var results []int
	results = append(results, a.getMaxModelNumber(input))
	results = append(results, a.getMinModelNumber(input))
	return results
}

func (a *arithmeticLogicUnit) getMaxModelNumber(input []string) int {
	return FindModelNumber(input, true)
}

func (a *arithmeticLogicUnit) getMinModelNumber(input []string) int {
	return FindModelNumber(input, false)
}

func NewArithmeticLogicUnit() Problem {
	return &arithmeticLogicUnit{
		dailyProblem{
			day: 24,
		},
	}
}
