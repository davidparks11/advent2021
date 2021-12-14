package advent

import (
	"strings"
)

var _ Problem = &sevenSegmentSearch{}

type sevenSegmentSearch struct {
	dailyProblem
}

func NewSevenSegmentSearch() Problem {
	return &sevenSegmentSearch{
		dailyProblem{
			day: 8,
		},
	}
}

func (s *sevenSegmentSearch) Solve() interface{} {
	input := s.GetInputLines()
	var results []int
	results = append(results, s.uniqueSum(input))
	results = append(results, s.outputSum(input))
	return results
}

func (s *sevenSegmentSearch) parseInput(input []string) []*entry {
	newSegmentFromString := func(str string) pattern {
		var segment pattern
		for _, char := range str {
			segment |= 1 << (char - 'a')
		}
		return segment
	}

	var signalPatterns []*entry
	for _, line := range input {
		pipeSplit := strings.Split(strings.TrimSpace(line), " | ")

		var patterns []pattern
		var output []pattern
		for _, patternString := range strings.Split(pipeSplit[0], " ") {
			patterns = append(patterns, newSegmentFromString(patternString))
		}
		for _, outputString := range strings.Split(pipeSplit[1], " ") {
			output = append(output, newSegmentFromString(outputString))
		}

		signalPatterns = append(signalPatterns, &entry{
			patterns: patterns,
			output:   output,
		})
	}

	return signalPatterns
}

/*
You barely reach the safety of the cave when the whale smashes into the cave mouth, collapsing it. Sensors indicate another exit to this cave at a much greater depth, so you have no choice but to press on.

As your submarine slowly makes its way through the cave system, you notice that the four-digit seven-segment displays in your submarine are malfunctioning; they must have been damaged during the escape. You'll be in a lot of trouble without them, so you'd better figure out what's wrong.

Each digit of a seven-segment display is rendered by turning on or off any of seven segments named a through g:

  0:      1:      2:      3:      4:
 aaaa    ....    aaaa    aaaa    ....
b    c  .    c  .    c  .    c  b    c
b    c  .    c  .    c  .    c  b    c
 ....    ....    dddd    dddd    dddd
e    f  .    f  e    .  .    f  .    f
e    f  .    f  e    .  .    f  .    f
 gggg    ....    gggg    gggg    ....

  5:      6:      7:      8:      9:
 aaaa    aaaa    aaaa    aaaa    aaaa
b    .  b    .  .    c  b    c  b    c
b    .  b    .  .    c  b    c  b    c
 dddd    dddd    ....    dddd    dddd
.    f  e    f  .    f  e    f  .    f
.    f  e    f  .    f  e    f  .    f
 gggg    gggg    ....    gggg    gggg
So, to render a 1, only segments c and f would be turned on; the rest would be off. To render a 7, only segments a, c, and f would be turned on.

The problem is that the signals which control the segments have been mixed up on each display. The submarine is still trying to display numbers by producing output on signal wires a through g, but those wires are connected to segments randomly. Worse, the wire/segment connections are mixed up separately for each four-digit display! (All of the digits within a display use the same connections, though.)

So, you might know that only signal wires b and g are turned on, but that doesn't mean segments b and g are turned on: the only digit that uses two segments is 1, so it must mean segments c and f are meant to be on. With just that information, you still can't tell which wire (b/g) goes to which segment (c/f). For that, you'll need to collect more information.

For each display, you watch the changing signals for a while, make a note of all ten unique signal patterns you see, and then write down a single four digit output value (your puzzle input). Using the signal patterns, you should be able to work out which pattern corresponds to which digit.

For example, here is what you might see in a single entry in your notes:

acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf
(The entry is wrapped here to two lines so it fits; in your notes, it will all be on a single line.)

Each entry consists of ten unique signal patterns, a | delimiter, and finally the four digit output value. Within an entry, the same wire/segment connections are used (but you don't know what the connections actually are). The unique signal patterns correspond to the ten different ways the submarine tries to render a digit using the current wire/segment connections. Because 7 is the only digit that uses three segments, dab in the above example means that to render a 7, signal lines d, a, and b are on. Because 4 is the only digit that uses four segments, eafb means that to render a 4, signal lines e, a, f, and b are on.

Using this information, you should be able to work out which combination of signal wires corresponds to each of the ten digits. Then, you can decode the four digit output value. Unfortunately, in the above example, all of the digits in the output value (cdfeb fcadb cdfeb cdbaf) use five segments and are more difficult to deduce.

For now, focus on the easy digits. Consider this larger example:

be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb |
fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec |
fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef |
cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega |
efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga |
gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf |
gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf |
cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd |
ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg |
gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc |
fgae cfgab fg bagce
Because the digits 1, 4, 7, and 8 each use a unique number of segments, you should be able to tell which combinations of signals correspond to those digits. Counting only digits in the output values (the part after | on each line), in the above example, there are 26 instances of digits that use a unique number of segments (highlighted above).

In the output values, how many times do digits 1, 4, 7, or 8 appear?
*/
func (s *sevenSegmentSearch) uniqueSum(input []string) int {
	signalPattens := s.parseInput(input)
	distinctSigCount := 0
	for _, p := range signalPattens {
		for _, o := range p.output {
			if o.oneCount() == 2 || o.oneCount() == 3 || o.oneCount() == 4 || o.oneCount() == 7 {
				distinctSigCount++
			}
		}
	}
	return distinctSigCount
}

/*
Through a little deduction, you should now be able to determine the remaining digits. Consider again the first example above:

acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab |
cdfeb fcadb cdfeb cdbaf
After some careful analysis, the mapping between signal wires and segments only make sense in the following configuration:

 dddd
e    a
e    a
 ffff
g    b
g    b
 cccc
So, the unique signal patterns would correspond to the following digits:

acedgfb: 8
cdfbe: 5
gcdfa: 2
fbcad: 3
dab: 7
cefabd: 9
cdfgeb: 6
eafb: 4
cagedb: 0
ab: 1
Then, the four digits of the output value can be decoded:

cdfeb: 5
fcadb: 3
cdfeb: 5
cdbaf: 3
Therefore, the output value for this entry is 5353.

Following this same process for each entry in the second, larger example above, the output value of each entry can be determined:

fdgacbe cefdb cefbgd gcbe: 8394
fcgedb cgb dgebacf gc: 9781
cg cg fdcagb cbg: 1197
efabcd cedba gadfec cb: 9361
gecf egdcabf bgf bfgea: 4873
gebdcfa ecba ca fadegcb: 8418
cefg dcbef fcge gbcadfe: 4548
ed bcgafe cdgba cbgef: 1625
gbdfcae bgc cg cgb: 8717
fgae cfgab fg bagce: 4315
Adding all of the output values in this larger example produces 61229.

For each entry, determine all of the wire/segment connections and decode the four-digit output values. What do you get if you add up all of the output values?
*/
func (s *sevenSegmentSearch) outputSum(input []string) int {
	entries := s.parseInput(input)
	sumOutputs := 0
	for _, e := range entries {
		var digits [10]pattern
		for _, p := range e.patterns {
			s.solveUniqueCases(&digits, p)
		}
		for _, p := range e.patterns {
			s.inferOtherCases(&digits, p)
		}
		decodeOutput := 0
		for _, o := range e.output {
			decodeOutput = decodeOutput*10 + s.matchDigit(&digits, o)
		}
		sumOutputs += decodeOutput
	}
	return sumOutputs
}

func (s *sevenSegmentSearch) matchDigit(digits *[10]pattern, p pattern) int {
	for i, d := range digits {
		if d == p {
			return i
		}
	}
	return -1
}

func (s *sevenSegmentSearch) solveUniqueCases(digits *[10]pattern, p pattern) {
	switch p.oneCount() {
	case 2:
		digits[1] = p
	case 3:
		digits[7] = p
	case 4:
		digits[4] = p
	case 7:
		digits[8] = p
	}
}

func (s *sevenSegmentSearch) inferOtherCases(digits *[10]pattern, p pattern) {
	//all non-unique cases have a one count of 5 or 6. Their number can be
	//found with the one count and unique cases
	switch p.oneCount() {
	case 5:
		if p|digits[4] == digits[8] {
			digits[2] = p
		} else if p&digits[7] == digits[7] {
			digits[3] = p
		} else {
			digits[5] = p
		}
	case 6:
		if p&digits[4] == digits[4] {
			digits[9] = p
		} else if p&digits[7] == digits[7] {
			digits[0] = p
		} else {
			digits[6] = p
		}
	}
}

type entry struct {
	patterns []pattern
	output   []pattern
}

//a pattern stores segment information in 7 bit bitmask in the order gfedcba.
//Ex: cfgab = 1100111
type pattern uint8

func (s *pattern) oneCount() uint8 {
	var count uint8
	for n := 0; n < 7; n++ {
		count += uint8(*s) >> n & 1
	}
	return count
}
