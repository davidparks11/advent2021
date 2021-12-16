package day5

import (
	"github.com/davidparks11/advent2021/internal/coordinate"
	"github.com/davidparks11/advent2021/internal/math"
	"strconv"
	"strings"
)

type Line struct {
	P1 coordinate.Point
	P2 coordinate.Point
}

func ParseInput(input []string, allowDiagonals bool) []*Line {
	var lines []*Line
	for _, inputLine := range input {
		l := parseLine(inputLine)
		if !(l.isHorizontal() || l.isVertical()) {
			if !(allowDiagonals && l.isDiagonal()) {
				continue
			}
		}
		lines = append(lines, l)
	}
	return lines
}

func parseLine(input string) *Line {
	var nums []int
	for _, numPair := range strings.Split(input, " -> ") {
		for _, numStr := range strings.Split(numPair, ",") {
			num, err := strconv.ParseInt(numStr, 10, 32)
			if err != nil {
				panic(err.Error())
			}
			nums = append(nums, int(num))
		}
	}

	return &Line{
		P1: coordinate.Point{X: nums[0], Y: nums[1]},
		P2: coordinate.Point{X: nums[2], Y: nums[3]},
	}
}

func (l *Line) isHorizontal() bool {
	return l.P1.Y == l.P2.Y
}

func (l *Line) isVertical() bool {
	return l.P1.X == l.P2.X
}

func (l *Line) isDiagonal() bool {
	return math.Abs(l.P1.X-l.P2.X) == math.Abs(l.P1.Y-l.P2.Y)
}

//returns Points along Line from P1 to P2 inclusively
func (l *Line) Points() []*coordinate.Point {
	var points []*coordinate.Point
	xDist := l.P2.X - l.P1.X
	yDist := l.P2.Y - l.P1.Y
	xSign := math.Sign(xDist)
	ySign := math.Sign(yDist)

	delta := math.Max(math.Abs(xDist), math.Abs(yDist))
	for i := 0; i <= delta; i++ {
		points = append(points, &coordinate.Point{
			X: l.P1.X + xSign*i,
			Y: l.P1.Y + ySign*i,
		})
	}

	return points
}
