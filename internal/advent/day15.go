package advent

import (
)

type chiton struct {
	dailyProblem
}

func NewChiton() Problem {
	return &chiton{
		dailyProblem{
			day: 15,
		},
	}
}

func (c *chiton) Solve() interface{} {
	input := c.GetInputLines()
	var results []int
	results = append(results, c.part1(input))
	return results
}

func (c *chiton) part1(input []string) int {
	// riskLevels := asciiNumGridToIntArray(input)
	return 0
}