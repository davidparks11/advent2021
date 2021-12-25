package advent

import (
	"fmt"

	"github.com/davidparks11/advent2021/internal/coordinate"
)

type seaCucumber struct {
	dailyProblem
}

func NewSeaCumber() Problem {
	return &seaCucumber{
		dailyProblem{
			day: 25,
		},
	}
}

func (s *seaCucumber) Solve() interface{} {
	input := s.GetInputLines()
	var results []int
	results = append(results, s.cucumberSteps(input))
	return results
}

/*
This is it: the bottom of the ocean trench, the last place the sleigh keys could be. Your submarine's experimental antenna still isn't boosted enough to detect the keys, but they must be here. All you need to do is reach the seafloor and find them.

At least, you'd touch down on the seafloor if you could; unfortunately, it's completely covered by two large herds of sea cucumbers, and there isn't an open space large enough for your submarine.

You suspect that the Elves must have done this before, because just then you discover the phone number of a deep-sea marine biologist on a handwritten note taped to the wall of the submarine's cockpit.

"Sea cucumbers? Yeah, they're probably hunting for food. But don't worry, they're predictable critters: they move in perfectly straight lines, only moving forward when there's space to do so. They're actually quite polite!"

You explain that you'd like to predict when you could land your submarine.

"Oh that's easy, they'll eventually pile up and leave enough space for-- wait, did you say submarine? And the only place with that many sea cucumbers would be at the very bottom of the Mariana--" You hang up the phone.

There are two herds of sea cucumbers sharing the same region; one always moves east (>), while the other always moves south (v). Each location can contain at most one sea cucumber; the remaining locations are empty (.). The submarine helpfully generates a map of the situation (your puzzle input). For example:

v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>
Every step, the sea cucumbers in the east-facing herd attempt to move forward one location, then the sea cucumbers in the south-facing herd attempt to move forward one location. When a herd moves forward, every sea cucumber in the herd first simultaneously considers whether there is a sea cucumber in the adjacent location it's facing (even another sea cucumber facing the same direction), and then every sea cucumber facing an empty location simultaneously moves into that location.

So, in a situation like this:

...>>>>>...
After one step, only the rightmost sea cucumber would have moved:

...>>>>.>..
After the next step, two sea cucumbers move:

...>>>.>.>.
During a single step, the east-facing herd moves first, then the south-facing herd moves. So, given this situation:

..........
.>v....v..
.......>..
..........
After a single step, of the sea cucumbers on the left, only the south-facing sea cucumber has moved (as it wasn't out of the way in time for the east-facing cucumber on the left to move), but both sea cucumbers on the right have moved (as the east-facing sea cucumber moved out of the way of the south-facing sea cucumber):

..........
.>........
..v....v>.
..........
Due to strong water currents in the area, sea cucumbers that move off the right edge of the map appear on the left edge, and sea cucumbers that move off the bottom edge of the map appear on the top edge. Sea cucumbers always check whether their destination location is empty before moving, even if that destination is on the opposite side of the map:

Initial state:
...>...
.......
......>
v.....>
......>
.......
..vvv..

After 1 step:
..vv>..
.......
>......
v.....>
>......
.......
....v..

After 2 steps:
....v>.
..vv...
.>.....
......>
v>.....
.......
.......

After 3 steps:
......>
..v.v..
..>v...
>......
..>....
v......
.......

After 4 steps:
>......
..v....
..>.v..
.>.v...
...>...
.......
v......
To find a safe place to land your submarine, the sea cucumbers need to stop moving. Again consider the first example:

Initial state:
v...>>.vv>
.vv>>.vv..
>>.>v>...v
>>v>>.>.v.
v>v.vv.v..
>.>>..v...
.vv..>.>v.
v.v..>>v.v
....v..v.>

...

After 58 steps:
..>>v>vv..
..v.>>vv..
..>>v>>vv.
..>>>>>vv.
v......>vv
v>v....>>v
vvv.....>>
>vv......>
.>v.vv.v..
In this example, the sea cucumbers stop moving after 58 steps.

Find somewhere safe to land your submarine. What is the first step on which no sea cucumbers move?
*/
func (s *seaCucumber) cucumberSteps(input []string) int {
	heard := s.parseInput(input)
	count := 1
	for ; !heard.step(); count++ {}
	return count
}

type cucumber rune

func (c cucumber) nextPosition() coordinate.Point {
	if c == south {
		return coordinate.Point{X: 0, Y: 1}
	}
	return coordinate.Point{X: 1, Y: 0}
}

const (
	south cucumber = 'v'
	east  cucumber = '>'
)

type cucumberHeard struct {
	length, width int
	cucumbers     map[coordinate.Point]cucumber
}

func (c *cucumberHeard) String() string {
	str := fmt.Sprintf("length: %d, width %d\n", c.length, c.width)
	for y := 0; y < c.length; y++ {
		for x := 0; x < c.width; x++ {
			if r, found := c.cucumbers[coordinate.Point{Y: y, X: x}]; found {
				str += string(r)
			} else {
				str += "."
			}
		}
		str += "\n"
	}
	return str
}

func (c *cucumberHeard) step() bool {
	done := true
	newPositions := make(map[coordinate.Point]cucumber)
	for p, current := range c.cucumbers {
		if current == south {
			continue
		}
		if nextPos := c.nextPosition(p); c.isFree(current, nextPos, nil) {
			newPositions[nextPos] = current
			done = false
		} else {
			newPositions[p] = current
		}
	}

	for p, current := range c.cucumbers {
		if current == east {
			continue
		}
		if nextPos := c.nextPosition(p); c.isFree(current, nextPos, newPositions) {
			newPositions[nextPos] = current
			done = false
		} else {
			newPositions[p] = current
		}
	}

	c.cucumbers = newPositions

	return done
}

func (c *cucumberHeard) nextPosition(p coordinate.Point) coordinate.Point {
	delta := c.cucumbers[p].nextPosition()
	next := coordinate.Point{X: p.X + delta.X, Y: p.Y + delta.Y}
	if next.Y == c.length {
		next.Y = 0
	}
	if next.X == c.width {
		next.X = 0
	}
	return next
}

func (c *cucumberHeard) isFree(current cucumber, p coordinate.Point, newPositions map[coordinate.Point]cucumber) bool {
	if current == east {
		if _, found := c.cucumbers[p]; found {
			return false
		}
	} else {
		if other, found := c.cucumbers[p]; found && other == south {
			return false
		}
		if other, found := newPositions[p]; found && other == east{
			return false
		}
	}
	return true
}

func (c *seaCucumber) parseInput(input []string) (heard *cucumberHeard) {
	heard = &cucumberHeard{
		length:    len(input),
		width:     len(input[0]),
		cucumbers: make(map[coordinate.Point]cucumber),
	}
	for y, line := range input {
		for x, char := range line {
			if char == rune(south) || char == rune(east) {
				heard.cucumbers[coordinate.Point{X: x, Y: y}] = cucumber(char)
			}
		}
	}
	return heard
}
