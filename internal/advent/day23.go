package advent

import (
	"strings"

	"github.com/davidparks11/advent2021/internal/coordinate"
	"github.com/davidparks11/advent2021/internal/math"
)

type amphipodProb struct {
	dailyProblem
}

func (a *amphipodProb) Solve() interface{} {
	input := a.GetInputLines()
	var results []int
	results = append(results, a.loswestFoldedSortCost(input))
	results = append(results, a.lowestSortCost(input))
	return results
}

func NewAmphipodProb() Problem {
	return &amphipodProb{
		dailyProblem{
			day: 23,
		},
	}
}

/*
A group of amphipods notice your fancy submarine and flag you down. "With such an impressive shell," one amphipod says, "surely you can help us with a question that has stumped our best scientists."

They go on to explain that a group of timid, stubborn amphipods live in a nearby burrow. Four types of amphipods live there: Amber (A), Bronze (B), Copper (C), and Desert (D). They live in a burrow that consists of a hallway and four side rooms. The side rooms are initially full of amphipods, and the hallway is initially empty.

They give you a diagram of the situation (your puzzle input), including locations of each amphipod (A, B, C, or D, each of which is occupying an otherwise open space), walls (#), and open space (.).

For example:

#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
The amphipods would like a method to organize every amphipod into side rooms so that each side room contains one type of amphipod and the types are sorted A-D going left to right, like this:

#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########
Amphipods can move up, down, left, or right so long as they are moving into an unoccupied open space. Each type of amphipod requires a different amount of energy to move one step: Amber amphipods require 1 energy per step, Bronze amphipods require 10 energy, Copper amphipods require 100, and Desert ones require 1000. The amphipods would like you to find a way to organize the amphipods that requires the least total energy.

However, because they are timid and stubborn, the amphipods have some extra rules:

Amphipods will never stop on the space immediately outside any room. They can move into that space so long as they immediately continue moving. (Specifically, this refers to the four open spaces in the hallway that are directly above an amphipod starting position.)
Amphipods will never move from the hallway into a room unless that room is their destination room and that room contains no amphipods which do not also have that room as their own destination. If an amphipod's starting room is not its destination room, it can stay in that room until it leaves the room. (For example, an Amber amphipod will not move from the hallway into the right three rooms, and will only move into the leftmost room if that room is empty or if it only contains other Amber amphipods.)
Once an amphipod stops moving in the hallway, it will stay in that spot until it can move into a room. (That is, once any amphipod starts moving, any other amphipods currently in the hallway are locked in place and will not move again until they can move fully into a room.)
In the above example, the amphipods can be organized using a minimum of 12521 energy. One way to do this is shown below.

Starting configuration:

#############
#...........#
###B#C#B#D###
  #A#D#C#A#
  #########
One Bronze amphipod moves into the hallway, taking 4 steps and using 40 energy:

#############
#...B.......#
###B#C#.#D###
  #A#D#C#A#
  #########
The only Copper amphipod not in its side room moves there, taking 4 steps and using 400 energy:

#############
#...B.......#
###B#.#C#D###
  #A#D#C#A#
  #########
A Desert amphipod moves out of the way, taking 3 steps and using 3000 energy, and then the Bronze amphipod takes its place, taking 3 steps and using 30 energy:

#############
#.....D.....#
###B#.#C#D###
  #A#B#C#A#
  #########
The leftmost Bronze amphipod moves to its room using 40 energy:

#############
#.....D.....#
###.#B#C#D###
  #A#B#C#A#
  #########
Both amphipods in the rightmost room move into the hallway, using 2003 energy in total:

#############
#.....D.D.A.#
###.#B#C#.###
  #A#B#C#.#
  #########
Both Desert amphipods move into the rightmost room using 7000 energy:

#############
#.........A.#
###.#B#C#D###
  #A#B#C#D#
  #########
Finally, the last Amber amphipod moves into its room, using 8 energy:

#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #########
What is the least energy required to organize the amphipods?
*/
func (a *amphipodProb) loswestFoldedSortCost(input []string) int {
	return organizeAmphipods(input, Burrow(targetFoldedBurrow))
}

/*
As you prepare to give the amphipods your solution, you notice that the diagram they handed you was actually folded up. As you unfold it, you discover an extra part of the diagram.

Between the first and second lines of text that contain amphipod starting positions, insert the following lines:

  #D#C#B#A#
  #D#B#A#C#
So, the above example now becomes:

#############
#...........#
###B#C#B#D###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########
The amphipods still want to be organized into rooms similar to before:

#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #A#B#C#D#
  #A#B#C#D#
  #########
In this updated example, the least energy required to organize these amphipods is 44169:

#############
#...........#
###B#C#B#D###
  #D#C#B#A#
  #D#B#A#C#
  #A#D#C#A#
  #########

...

#############
#...........#
###A#B#C#D###
  #A#B#C#D#
  #A#B#C#D#
  #A#B#C#D#
  #########
Using the initial configuration from the full diagram, what is the least energy required to organize the amphipods?
*/
func (a *amphipodProb) lowestSortCost(input []string) int {
	input = []string{input[0], input[1], input[2], "  #D#C#B#A#  ", "  #D#B#A#C#  ", input[3], input[4]}
	return organizeAmphipods(input, Burrow(targetBurrow))
}

var (
	roomA   = []coordinate.Point{{Y: 2, X: 3}, {Y: 3, X: 3}, {Y: 4, X: 3}, {Y: 5, X: 3}}
	roomB   = []coordinate.Point{{Y: 2, X: 5}, {Y: 3, X: 5}, {Y: 4, X: 5}, {Y: 5, X: 5}}
	roomC   = []coordinate.Point{{Y: 2, X: 7}, {Y: 3, X: 7}, {Y: 4, X: 7}, {Y: 5, X: 7}}
	roomD   = []coordinate.Point{{Y: 2, X: 9}, {Y: 3, X: 9}, {Y: 4, X: 9}, {Y: 5, X: 9}}
	rooms   = [][]coordinate.Point{roomA, roomB, roomC, roomD}
	hallway = []coordinate.Point{{Y: 1, X: 1}, {Y: 1, X: 2}, {Y: 1, X: 4}, {Y: 1, X: 6}, {Y: 1, X: 8}, {Y: 1, X: 10}, {Y: 1, X: 11}}
)

const (
	amber  uint8 = 'A'
	bronze uint8 = 'B'
	copper uint8 = 'C'
	desert uint8 = 'D'
)

var costByAmphipod = map[uint8]int{
	amber:  1,
	bronze: 10,
	copper: 100,
	desert: 1000,
}

var roomsByeAmphipod = map[uint8][]coordinate.Point{
	amber:  roomA,
	bronze: roomB,
	copper: roomC,
	desert: roomD,
}


type state struct {
	burrow Burrow
	cost   int
}

type Burrow string


func organizeAmphipods(input []string, target Burrow) int {
	start := Burrow(strings.Join(input, "\n"))

	costByStates := map[Burrow]int{start: 0}

	queue := []Burrow{start}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, step := range current.moves() {
			if _, ok := costByStates[step.burrow]; !ok {
				costByStates[step.burrow] = costByStates[current] + step.cost
				queue = append(queue, step.burrow)
			} else {
				if costByStates[step.burrow] > costByStates[current]+step.cost {
					costByStates[step.burrow] = costByStates[current] + step.cost
					queue = append(queue, step.burrow)
				}
			}
		}
	}

	return costByStates[target]
}

func (burrow Burrow) moves() []state {
	
	roomLength := 4
	if burrow.isFolded() {
		roomLength = 2
	}
	
	var states []state
	for _, room := range rooms {
		for _, rPos := range room[:roomLength] {
			if !burrow.isEmpty(rPos) {
				states = append(states, burrow.moveToHallway(rPos)...)
				break
			}
		}
	}

	for _, hPos := range hallway {
		if !burrow.isEmpty(hPos) {
			states = append(states, burrow.moveToRoom(hPos)...)
		}
	}

	return states
}

func (burrow Burrow) moveToHallway(current coordinate.Point) []state {
	room, found := roomsByeAmphipod[burrow[loc(current)]]
	if !found {
		panic("room does not exist")
	}

	pathClear := false
	if burrow.isFolded() {
		room = room[:2]
	}
	if room[0].X == current.X {
		for _, r := range room {
			if !burrow.isEmpty(r) && burrow[loc(r)] != burrow[loc(current)] {
				pathClear = true
				break
			}
		}
	} else {
		pathClear = true
	}

	if !pathClear {
		return nil
	}

	var possiblePositions []coordinate.Point
	for _, hPos := range hallway {
		if burrow.isEmpty(hPos) {
			possiblePositions = append(possiblePositions, hPos)
		} else if hPos.X < current.X {
			possiblePositions = possiblePositions[:0]
		} else {
			break
		}
	}

	var states []state
	for _, pos := range possiblePositions {
		next := Burrow(burrow[:loc(pos)] + Burrow(burrow[loc(current)]) + burrow[loc(pos)+1:loc(current)] + "." + burrow[loc(current)+1:])
		energy := (math.Abs(current.X-pos.X) + math.Abs(current.Y-pos.Y)) * costByAmphipod[burrow[loc(current)]]
		states = append(states, state{next, energy})
	}

	return states
}

func (burrow Burrow) moveToRoom(a coordinate.Point) []state {
	room, found := roomsByeAmphipod[burrow[loc(a)]]
	if !found {
		panic("room does not exist")
	}

	if a.X < room[0].X {
		for _, hPos := range hallway {
			if hPos.X <= a.X {
				continue
			}
			if hPos.X > room[0].X {
				break
			}
			if !burrow.isEmpty(hPos) {
				return nil
			}
		}
	} else {
		for i := len(hallway) - 1; i >= 0; i-- {
			hPos := hallway[i]
			if hPos.X >= a.X {
				continue
			}
			if hPos.X < room[0].X {
				break
			}
			if !burrow.isEmpty(hPos) {
				return nil
			}
		}
	}

	if burrow.isFolded() {
		room = room[:2]
	}

	for _, r := range room {
		if !burrow.isEmpty(r) && burrow[loc(r)] != burrow[loc(a)] {
			return nil
		}
	}

	var states []state
	for i := len(room) - 1; i >= 0; i-- {
		r := room[i]
		if burrow.isEmpty(r) {
			next := Burrow(burrow[:loc(a)] + "." + burrow[loc(a)+1:loc(r)] + Burrow(burrow[loc(a)]) + burrow[loc(r)+1:])
			e := (math.Abs(a.X-r.X) + math.Abs(a.Y-r.Y)) * costByAmphipod[burrow[loc(a)]]
			states = append(states, state{next, e})
			break
		}
	}

	return states
}

func (burrow Burrow) isFolded() bool {
	return len(burrow)/14 <= 5
}

func (burrow Burrow) isEmpty(pos coordinate.Point) bool {
	return burrow[loc(pos)] == '.'
}

func loc(p coordinate.Point) int {
	return p.Y*14 + p.X
}

var targetFoldedBurrow = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#  
  #########  `

var targetBurrow = `#############
#...........#
###A#B#C#D###
  #A#B#C#D#  
  #A#B#C#D#  
  #A#B#C#D#  
  #########  `
