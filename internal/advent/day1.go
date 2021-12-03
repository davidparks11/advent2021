package advent

import (
	"strconv"
)

var _ Problem = &sonorSweep{}

type sonorSweep struct {
	dailyProblem
}

func NewSonorSweep(day int) *sonorSweep {
	return &sonorSweep{
		dailyProblem{day: day},
	}
}

func (r *sonorSweep) Solve() []string {
	input := IntsFromStrings(r.GetInputLines())
	var results []string
	results = append(results, strconv.Itoa(countDepthIncreases(input)))
	results = append(results, strconv.Itoa(count3WideDepthIncreases(input)))

	return results
}

// The first order of business is to figure out how quickly the depth increases,
// just so you know what you're dealing with - you never know if the keys will
// get carried into deeper water by an ocean current or a fish or something.
// To do this, count the number of times a depth measurement increases from the
// previous measurement. (There is no measurement before the first measurement.)
func countDepthIncreases(input []int) (increases int) {
	if len(input) < 2 {
		return //can't have an increase with one or zero elements
	}

	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			increases++
		}
	}

	return
}

// Your goal now is to count the number of times the sum of measurements in this
// sliding window increases from the previous sum. So, compare A with B, then
// compare B with C, then C with D, and so on. Stop when there aren't enough
// measurements left to create a new three-measurement sum.
func count3WideDepthIncreases(input []int) (increases int) {
	if len(input) < 4 {
		return //similar situation as above
	}

	summedDepth := input[0] + input[1] + input[2]
	nextSummedDepth := 0
	for i := 3; i < len(input); i++ {
		nextSummedDepth = summedDepth + input[i] - input[i-3]
		if nextSummedDepth > summedDepth {
			increases++
		}
		summedDepth = nextSummedDepth
	}

	return
}
