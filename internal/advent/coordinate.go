package advent

import "github.com/davidparks11/advent2021/internal/math"

type point struct {
	x int
	y int
}

type line struct {
	p1 point
	p2 point
}

func (l *line) isHorizontal() bool {
	return l.p1.y == l.p2.y
}

func (l *line) isVertical() bool {
	return l.p1.x == l.p2.x
}

func (l *line) isDiagonal() bool {
	return math.Abs(l.p1.x-l.p2.x) == math.Abs(l.p1.y-l.p2.y)
}

//returns points along line from p1 to p2 inclusively
func (l *line) points() []*point {
	var points []*point
	xDist := l.p2.x - l.p1.x
	yDist := l.p2.y - l.p1.y
	xSign := math.Sign(xDist)
	ySign := math.Sign(yDist)

	delta := math.Max(math.Abs(xDist), math.Abs(yDist))
	for i := 0; i <= delta; i++ {
		points = append(points, &point{
			x: l.p1.x + xSign*i,
			y: l.p1.y + ySign*i,
		})
	}

	return points
}
