package day8

import "strings"

type Entry struct {
	Patterns []Pattern
	Outputs  []Pattern
}

func ParseInput(input []string) []*Entry {
	newSegmentFromString := func(str string) Pattern {
		var segment Pattern
		for _, char := range str {
			segment |= 1 << (char - 'a')
		}
		return segment
	}

	var signalPatterns []*Entry
	for _, line := range input {
		pipeSplit := strings.Split(strings.TrimSpace(line), " | ")

		var patterns []Pattern
		var output []Pattern
		for _, patternString := range strings.Split(pipeSplit[0], " ") {
			patterns = append(patterns, newSegmentFromString(patternString))
		}
		for _, outputString := range strings.Split(pipeSplit[1], " ") {
			output = append(output, newSegmentFromString(outputString))
		}

		signalPatterns = append(signalPatterns, &Entry{
			Patterns: patterns,
			Outputs:  output,
		})
	}

	return signalPatterns
}

//A Pattern stores segment information in 7 bit bitmask in the order gfedcba.
//Ex: cfgab = 1100111
type Pattern uint8

func (s *Pattern) OneCount() uint8 {
	var count uint8
	for n := 0; n < 7; n++ {
		count += uint8(*s) >> n & 1
	}
	return count
}

