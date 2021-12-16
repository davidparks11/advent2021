package advent

import "github.com/davidparks11/advent2021/internal/coordinate"

var _ Problem = &smokeBasin{}

type smokeBasin struct {
	dailyProblem
}

func NewSmokeBasin() Problem {
	return &smokeBasin{
		dailyProblem{
			day: 9,
		},
	}
}

func (s *smokeBasin) Solve() interface{} {
	input := asciiNumGridToIntArray(s.GetInputLines())
	var results []int
	results = append(results, s.sumRiskLevels(input))
	results = append(results, s.threeLargestBasins(input))
	return results
}

/*
These caves seem to be lava tubes. Parts are even still volcanically active; small hydrothermal vents release smoke into the caves that slowly settles like rain.

If you can model how the smoke flows through the caves, you might be able to avoid it and be that much safer. The submarine generates a heightmap of the floor of the nearby caves for you (your puzzle input).

Smoke flows to the lowest Point of the area it's in. For example, consider the following heightmap:

2199943210
3987894921
9856789892
8767896789
9899965678
Each number corresponds to the height of a particular location, where 9 is the highest and 0 is the lowest a location can be.

Your first goal is to find the low points - the locations that are lower than any of its adjacent locations. Most locations have four adjacent locations (up, down, left, and right); locations on the edge or corner of the map have three or two adjacent locations, respectively. (Diagonal locations do not count as adjacent.)

In the above example, there are four low points, all highlighted: two are in the first row (a 1 and a 0), one is in the third row (a 5), and one is in the bottom row (also a 5). All other locations on the heightmap have some lower adjacent location, and so are not low points.

The risk level of a low Point is 1 plus its height. In the above example, the risk levels of the low points are 2, 1, 6, and 6. The sum of the risk levels of all low points in the heightmap is therefore 15.

Find all of the low points on your heightmap. What is the sum of the risk levels of all low points on your heightmap?
*/
func (s *smokeBasin) sumRiskLevels(locationHeights [][]int) int {
	//trivial solution
	riskLevels := 0
	for y := 0; y < len(locationHeights); y++ {
		for x := 0; x < len(locationHeights[y]); x++ {
			if x > 0 && locationHeights[y][x] >= locationHeights[y][x-1] { //left location
				continue
			} else if x < len(locationHeights[y])-1 && locationHeights[y][x] >= locationHeights[y][x+1] { //right location
				continue
			} else if y > 0 && locationHeights[y][x] >= locationHeights[y-1][x] { //up location
				continue
			} else if y < len(locationHeights)-1 && locationHeights[y][x] >= locationHeights[y+1][x] { //down location
				continue
			}
			riskLevels += locationHeights[y][x] + 1
		}
	}
	return riskLevels
}

/*
Next, you need to find the largest basins so you know what areas are most important to avoid.

A basin is all locations that eventually flow downward to a single low Point. Therefore, every low Point has a basin, although some basins are very small. Locations of height 9 do not count as being in any basin, and all other locations will always be part of exactly one basin.

The size of a basin is the number of locations within the basin, including the low Point. The example above has four basins.

The top-left basin, size 3:

2199943210
3987894921
9856789892
8767896789
9899965678
The top-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
The middle basin, size 14:

2199943210
3987894921
9856789892
8767896789
9899965678
The bottom-right basin, size 9:

2199943210
3987894921
9856789892
8767896789
9899965678
Find the three largest basins and multiply their sizes together. In the above example, this is 9 * 14 * 9 = 1134.

What do you get if you multiply together the sizes of the three largest basins?
*/
func (s *smokeBasin) threeLargestBasins(locationHeights [][]int) int {

	lowPoints := []coordinate.Point{}
	for y := 0; y < len(locationHeights); y++ {
		for x := 0; x < len(locationHeights[y]); x++ {
			if isLowPoint := s.isLowestPoint(locationHeights, x, y); isLowPoint {
				lowPoints = append(lowPoints, coordinate.Point{x, y})
			}
		}
	}

	first, second, third := 0, 0, 0
	checkLargest := func(size int) {
		if size > first {
			first, second, third = size, first, second
		} else if size > second {
			second, third = size, second
		} else if size > third {
			third = size
		}
	}

	for _, p := range lowPoints {
		seen := make(map[coordinate.Point]struct{}) //set of points
		checkLargest(s.calcBasinSize(locationHeights, seen, p.X, p.Y))
	}

	return first * second * third
}

//counts number of locations in basin
func (s *smokeBasin) calcBasinSize(locationHeights [][]int, seen map[coordinate.Point]struct{}, x, y int) int {
	seen[coordinate.Point{X: x, Y: y}] = struct{}{}
	count := 1
	if _, found := seen[coordinate.Point{X: x - 1, Y: y}]; !found && s.inBounds(locationHeights, x-1, y) && locationHeights[y][x] < locationHeights[y][x-1] && locationHeights[y][x-1] != 9 {
		count += s.calcBasinSize(locationHeights, seen, x-1, y)
	}
	if _, found := seen[coordinate.Point{X: x + 1, Y: y}]; !found && s.inBounds(locationHeights, x+1, y) && locationHeights[y][x] < locationHeights[y][x+1] && locationHeights[y][x+1] != 9 {
		count += s.calcBasinSize(locationHeights, seen, x+1, y)
	}
	if _, found := seen[coordinate.Point{X: x, Y: y - 1}]; !found && s.inBounds(locationHeights, x, y-1) && locationHeights[y][x] < locationHeights[y-1][x] && locationHeights[y-1][x] != 9 {
		count += s.calcBasinSize(locationHeights, seen, x, y-1)
	}
	if _, found := seen[coordinate.Point{X: x, Y: y + 1}]; !found && s.inBounds(locationHeights, x, y+1) && locationHeights[y][x] < locationHeights[y+1][x] && locationHeights[y+1][x] != 9 {
		count += s.calcBasinSize(locationHeights, seen, x, y+1)
	}

	return count
}

//bounds check - assumes grid if not empty
func (s *smokeBasin) inBounds(grid [][]int, x, y int) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

//checks whether or not surrounding locations are less than current location
func (s *smokeBasin) isLowestPoint(grid [][]int, x, y int) bool {
	return (x > 0 && grid[y][x] < grid[y][x-1]) || //left location
		(x < len(grid[y])-1 && grid[y][x] < grid[y][x+1]) || //right location
		(y > 0 && grid[y][x] < grid[y-1][x]) || //up location
		(y < len(grid)-1 && grid[y][x] < grid[y+1][x]) //down location
}
