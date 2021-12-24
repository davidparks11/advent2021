package day22

import (
	"strconv"
	"strings"

	"github.com/davidparks11/advent2021/internal/math"
)

//helper to avoid code dup
type numRange struct {
	min, max int
}

func (c numRange) length() int {
	return math.Abs(c.max - c.min + 1)
}

func (c numRange) intersection(other numRange) *numRange {
	if other.max < c.min || c.max < other.min {
		return nil
	}

	return &numRange{min: math.Max(c.min, other.min), max: math.Min(c.max, other.max)}
}

type coordinateSet struct {
	On      bool
	x, y, z numRange
}

func NewCoordinateSetCube(min, max int, on bool) *coordinateSet {
	return &coordinateSet{
		On: on,
		x:  numRange{min: min, max: max},
		y:  numRange{min: min, max: max},
		z:  numRange{min: min, max: max},
	}
}

func (c coordinateSet) volume() int {
	return c.x.length() * c.y.length() * c.z.length()
}

func (c coordinateSet) intersection(other coordinateSet) *coordinateSet {
	if x, y, z := c.x.intersection(other.x), c.y.intersection(other.y), c.z.intersection(other.z); x != nil && y != nil && z != nil {
		return &coordinateSet{
			On: c.On,
			x:  *x,
			y:  *y,
			z:  *z,
		}
	}

	return nil
}

func (c coordinateSet) CountOnCubes(cubes []coordinateSet) int {
	//gather cubes that intersect with c
	var intersecting []coordinateSet
	for _, other := range cubes {
		if iCube := c.intersection(other); iCube != nil {
			intersecting = append(intersecting, *iCube)
		}
	}

	//find only volume of this cube minus volume of intersecting cube recursively
	volume := c.volume()
	for i, other := range intersecting {
		volume -= other.CountOnCubes(intersecting[i+1:])
	}

	return volume
}

//This is rough, but it works
func ParseInput(input []string, filterBounds *coordinateSet) []coordinateSet {
	parseLine := func(str string) *coordinateSet {
		var on bool
		if str[:2] == "on" {
			on = true
			str = str[3:]
		} else {
			str = str[4:]
		}

		var bounds []int
		boundsStrings := strings.Split(str, ",")
		for _, bString := range boundsStrings {

			numStrings := strings.Split(bString[2:], "..")
			min, err := strconv.Atoi(numStrings[0])
			if err != nil {
				panic(err)
			}

			max, err := strconv.Atoi(numStrings[1])
			if err != nil {
				panic(err)
			}

			if min > max {
				min, max = max, min
			}

			bounds = append(bounds, min, max)
		}

		c := &coordinateSet{
			On: on,
			x:  numRange{min: bounds[0], max: bounds[1]},
			y:  numRange{min: bounds[2], max: bounds[3]},
			z:  numRange{min: bounds[4], max: bounds[5]},
		}

		if filterBounds != nil {
			c = c.intersection(*filterBounds)
		}

		return c
	}

	var cubes []coordinateSet
	for _, line := range input {
		c := parseLine(line)
		if c != nil {
			cubes = append(cubes, *c)
		}
	}

	return cubes
}
