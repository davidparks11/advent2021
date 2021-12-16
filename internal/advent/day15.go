package advent

import (
	"container/heap"
	"math"

	. "github.com/davidparks11/advent2021/internal/advent/day15"
	"github.com/davidparks11/advent2021/internal/coordinate"
)

type chiton struct {
	dailyProblem
}

func NewChiton() Problem {
	return &chiton{
		dailyProblem{
			day: 15,
		},
	}
}

func (c *chiton) Solve() interface{} {
	input := c.GetInputLines()
	var results []int
	results = append(results, c.findLowestRisk(input))
	results = append(results, c.findExpandedLowestRisk(input))
	return results
}

/*
You've almost reached the exit of the cave, but the walls are getting closer together. Your submarine can barely still fit, though; the main problem is that the walls of the cave are covered in chitons, and it would be best not to bump any of them.

The cavern is large, but has a very low ceiling, restricting your motion to two dimensions. The shape of the cavern resembles a square; a quick scan of chiton density produces a map of risk level throughout the cave (your puzzle input). For example:

1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
You start in the top left position, your destination is the bottom right position, and you cannot move diagonally. The number at each position is its risk level; to determine the total risk of an entire path, add up the risk levels of each position you enter (that is, don't count the risk level of your starting position unless you enter it; leaving it adds no risk to your total).

Your goal is to find a path with the lowest total risk. In this example, a path with the lowest total risk is highlighted here:

1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581
The total risk of this path is 40 (the starting position is never entered, so its risk is not counted).

What is the lowest total risk of any path from the top left to the bottom right?
*/
func (c *chiton) findLowestRisk(input []string) int {
	riskLevels := Grid(asciiNumGridToIntArray(input))
	return c.dijkstra(riskLevels, coordinate.Point{}, coordinate.Point{X: riskLevels.Width() - 1, Y: riskLevels.Length() - 1})
}

/*
Now that you know how to find low-risk paths in the cave, you can try to find your way out.

The entire cave is actually five times larger in both dimensions than you thought; the area you originally scanned is just one tile in a 5x5 tile area that forms the full map. Your original map tile repeats to the right and downward; each time the tile repeats to the right or downward, all of its risk levels are 1 higher than the tile immediately up or left of it. However, risk levels above 9 wrap back around to 1. So, if your original map had some position with a risk level of 8, then that same position on each of the 25 total tiles would be as follows:

8 9 1 2 3
9 1 2 3 4
1 2 3 4 5
2 3 4 5 6
3 4 5 6 7
Each single digit above corresponds to the example position with a value of 8 on the top-left tile. Because the full map is actually five times larger in both dimensions, that position appears a total of 25 times, once in each duplicated tile, with the values shown above.

Here is the full five-times-as-large version of the first example above, with the original map in the top left corner highlighted:
...

Equipped with the full map, you can now find a path from the top left corner to the bottom right corner with the lowest total risk:

...

The total risk of this path is 315 (the starting position is still never entered, so its risk is not counted).

Using the full map, what is the lowest total risk of any path from the top left to the bottom right?
*/
func (c *chiton) findExpandedLowestRisk(input []string) int {
	riskLevels := Grid(asciiNumGridToIntArray(input)).Expand(5)
	return c.dijkstra(riskLevels, coordinate.Point{}, coordinate.Point{X: riskLevels.Width() - 1, Y: riskLevels.Length() - 1})
}

func (c *chiton) dijkstra(riskLevels Grid, start coordinate.Point, target coordinate.Point) int {

	distances := make(map[coordinate.Point]int)
	for y := 0; y < len(riskLevels); y++ {
		for x := 0; x < len(riskLevels[0]); x++ {
			distances[coordinate.Point{X: x, Y: y}] = math.MaxInt32
		}
	}

	distances[coordinate.Point{X: start.X, Y: start.Y}] = 0

	minHeap := &NodeHeap{&Node{Point: coordinate.Point{X: start.X, Y: start.Y}, Cost: 0}}
	heap.Init(minHeap)

	pointDeltas := []coordinate.Point{
		{X: 0, Y: -1}, {X: -1, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1},
	}

	for minHeap.Len() != 0 {
		current := heap.Pop(minHeap).(*Node)
		if current.Point == target {
			return current.Cost
		}

		if current.Cost > distances[current.Point] {
			continue
		}

		for _, delta := range pointDeltas {
			neighbor := coordinate.Point{X: current.X + delta.X, Y: current.Y + delta.Y}
			if riskLevels.InBounds(neighbor) {
				cost := current.Cost + riskLevels[neighbor.Y][neighbor.X]
				if cost < distances[neighbor] {
					heap.Push(minHeap, &Node{
						Point: neighbor,
						Cost:  cost,
					})
					distances[neighbor] = cost
				}
			}
		}
	}

	return 0
}
