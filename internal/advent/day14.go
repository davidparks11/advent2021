package advent

import (
	"math"
	"strings"
)

var _ Problem = &extendedPolymerization{}

type extendedPolymerization struct {
	dailyProblem
}

func NewExtendedPolymerization() Problem {
	return &extendedPolymerization{
		dailyProblem{
			day: 14,
		},
	}
}

func (e *extendedPolymerization) Solve() interface{} {
	input := e.GetInputLines()
	var results []int64
	results = append(results, e.minMax10(input))
	results = append(results, e.minMax40(input))
	return results
}

/*
The incredible pressures at this depth are starting to put a strain on your submarine. The submarine has polymerization equipment that would produce suitable materials to reinforce the submarine, and the nearby volcanically-active caves should even have the necessary input elements in sufficient quantities.

The submarine manual contains instructions for finding the optimal polymer formula; specifically, it offers a polymer template and a list of pair insertion rules (your puzzle input). You just need to work out what polymer would result after repeating the pair insertion process a few times.

For example:

NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C
The first line is the polymer template - this is the starting point of the process.

The following section defines the pair insertion rules. A rule like AB -> C means that when elements A and B are immediately adjacent, element C should be inserted between them. These insertions all happen simultaneously.

So, starting with the polymer template NNCB, the first step simultaneously considers all three pairs:

The first pair (NN) matches the rule NN -> C, so element C is inserted between the first N and the second N.
The second pair (NC) matches the rule NC -> B, so element B is inserted between the N and the C.
The third pair (CB) matches the rule CB -> H, so element H is inserted between the C and the B.
Note that these pairs overlap: the second element of one pair is the first element of the next pair. Also, because all pairs are considered simultaneously, inserted elements are not considered to be part of a pair until the next step.

After the first step of this process, the polymer becomes NCNBCHB.

Here are the results of a few steps using the above rules:

Template:     NNCB
After step 1: NCNBCHB
After step 2: NBCCNBBBCBHCB
After step 3: NBBBCNCCNBBNBNBBCHBHHBCHB
After step 4: NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB
This polymer grows quickly. After step 5, it has length 97; After step 10, it has length 3073. After step 10, B occurs 1749 times, C occurs 298 times, H occurs 161 times, and N occurs 865 times; taking the quantity of the most common element (B, 1749) and subtracting the quantity of the least common element (H, 161) produces 1749 - 161 = 1588.

Apply 10 steps of pair insertion to the polymer template and find the most and least common elements in the result. What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?

Your puzzle answer was 2967.
*/
func (e *extendedPolymerization) minMax10(input []string) int64 {
	template, insertions := e.parseInput(input)
	return e.minMaxDifference(template, insertions, 10)
}

/*

The resulting polymer isn't nearly strong enough to reinforce the submarine. You'll need to run more steps of the pair insertion process; a total of 40 steps should do it.

In the above example, the most common element is B (occurring 2192039569602 times) and the least common element is H (occurring 3849876073 times); subtracting these produces 2188189693529.

Apply 40 steps of pair insertion to the polymer template and find the most and least common elements in the result. What do you get if you take the quantity of the most common element and subtract the quantity of the least common element?
*/
func (e *extendedPolymerization) minMax40(input []string) int64 {
	template, insertions := e.parseInput(input)
	return e.minMaxDifference(template, insertions, 40)
}

func (e *extendedPolymerization) minMaxDifference(template string, insertions map[string]string, iterations int) int64 {
	pairCount := e.addTemplateCounts(template)

	for i := 0; i < iterations; i++ {
		pairCount = e.growPolymer(pairCount, insertions)
	}

	charCount := e.charCount(pairCount, template[len(template)-1])

	min, max := e.minMax(charCount)

	return max - min
}

func (e *extendedPolymerization) addTemplateCounts(template string) map[string]int64 {
	pairCount := make(map[string]int64)

	for i := 0; i < len(template)-1; i++ {
		pair := string(template[i]) + string(template[i+1])
		pairCount[pair]++
	}

	return pairCount
}

func (e *extendedPolymerization) growPolymer(pairCount map[string]int64, insertionRules map[string]string) map[string]int64 {
	nextCount := make(map[string]int64)

	for pair, count := range pairCount {
		rule := insertionRules[pair]
		nextCount[pair[0:1]+rule] += count
		nextCount[rule+pair[1:]] += count
	}

	return nextCount
}

func (e *extendedPolymerization) charCount(pairCount map[string]int64, lastChar byte) map[byte]int64 {
	characterCount := make(map[byte]int64)

	for pair, count := range pairCount {
		characterCount[pair[0]] += count
	}

	characterCount[lastChar]++ //correct off by one

	return characterCount
}

func (e *extendedPolymerization) minMax(count map[byte]int64) (int64, int64) {
	var min int64 = math.MaxInt64
	var max int64 = math.MinInt64

	for c := range count {
		if count[c] < min {
			min = count[c]
		}
		if count[c] > max {
			max = count[c]
		}
	}

	return min, max
}

func (e *extendedPolymerization) parseInput(input []string) (template string, insertions map[string]string) {
	insertions = make(map[string]string)
	template = input[0]
	for _, line := range input[2:] {
		insertionStrings := strings.Split(line, " -> ")
		insertions[insertionStrings[0]] = insertionStrings[1]
	}

	return
}
