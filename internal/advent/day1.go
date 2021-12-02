package advent

import (
	"strconv"
)

var _ Problem = &SonorSweep{}

type SonorSweep struct {
	dailyProblem
}

func (r *SonorSweep) Solve() {
	r.day = 1
	r.name = "Sonor Sweep"
	input := IntsFromStrings(r.GetInputLines())
	var results []string
	results = append(results, strconv.Itoa(countDepthIncreases(input)))
	r.WriteResult(results)
}

func countDepthIncreases(input []int) int {
	panic("not implemented")
}
