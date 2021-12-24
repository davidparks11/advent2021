package advent

import (
	. "github.com/davidparks11/advent2021/internal/advent/day22"
)

type reactorReboot struct {
	dailyProblem
}

func NewReactorReboot() Problem {
	return &reactorReboot{
		dailyProblem{
			day: 22,
		},
	}
}

func (d *reactorReboot) Solve() interface{} {
	input := d.GetInputLines()
	var results []int
	results = append(results, d.countOnWithinBounds(input))
	results = append(results, d.countOn(input))
	return results
}

func (d *reactorReboot) countOnWithinBounds(input []string) int {
	cubes := ParseInput(input, NewCoordinateSetCube(-50, 50, false))
	onCount := 0
	for i, c := range cubes {
		if c.On {
			o := c.CountOnCubes(cubes[i+1:])
			onCount += o
		}
	}
	return onCount
}

func (d *reactorReboot) countOn(input []string) int {
	cubes := ParseInput(input, nil)
	onCount := 0
	for i, c := range cubes {
		if c.On {
			onCount += c.CountOnCubes(cubes[i+1:])
		}
	}
	return onCount
}
