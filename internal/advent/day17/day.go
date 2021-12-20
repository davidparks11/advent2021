package day17

import "github.com/davidparks11/advent2021/internal/coordinate"

type Area struct {
	YMin int
	YMax int
	XMin int
	XMax int
}

//returns whether or not a point.X is greater than area.xMax or point.Y is less than area.yMin
func (a Area) Beyond(point coordinate.Point) bool {
	return point.X > a.XMax || point.Y < a.YMin
}

func (a *Area) Contains(p coordinate.Point) bool {
	return a.XMin <= p.X && p.X <= a.XMax && a.YMin <= p.Y && p.Y <= a.YMax
}
