package advent

var _ Problem = &dumboOctopus{}

type dumboOctopus struct {
	dailyProblem
}

func NewDumboOctopus() Problem {
	return &dumboOctopus{
		dailyProblem{
			day: 11,
		},
	}
}

func (d *dumboOctopus) Solve() []int {
	input := d.GetInputLines()
	var results []int
	results = append(results, d.flashCount(input))
	results = append(results, d.syncStep(input))

	return results
}

/*
You enter a large cavern full of rare bioluminescent dumbo octopuses! They seem to not like the Christmas lights on your submarine, so you turn them off for now.

There are 100 octopuses arranged neatly in a 10 by 10 grid. Each octopus slowly gains energy over time and flashes brightly for a moment when its energy is full. Although your lights are off, maybe you could navigate through the cave without disturbing the octopuses if you could predict when the flashes of light will happen.

Each octopus has an energy level - your submarine can remotely measure the energy level of each octopus (your puzzle input). For example:

5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
The energy level of each octopus is a value between 0 and 9. Here, the top-left octopus has an energy level of 5, the bottom-right one has an energy level of 6, and so on.

You can model the energy levels and flashes of light in steps. During a single step, the following occurs:

First, the energy level of each octopus increases by 1.
Then, any octopus with an energy level greater than 9 flashes. This increases the energy level of all adjacent octopuses by 1, including octopuses that are diagonally adjacent. If this causes an octopus to have an energy level greater than 9, it also flashes. This process continues as long as new octopuses keep having their energy level increased beyond 9. (An octopus can only flash at most once per step.)
Finally, any octopus that flashed during this step has its energy level set to 0, as it used all of its energy to flash.
Adjacent flashes can cause an octopus to flash on a step even if it begins that step with very little energy. Consider the middle octopus with 1 energy in this situation:

Before any steps:
11111
19991
19191
19991
11111

After step 1:
34543
40004
50005
40004
34543

After step 2:
45654
51115
61116
51115
45654
An octopus is highlighted when it flashed during the given step.

Here is how the larger example above progresses:

Before any steps:
.
.
.
After 100 steps, there have been a total of 1656 flashes.

Given the starting energy levels of the dumbo octopuses in your cavern, simulate 100 steps. How many total flashes are there after 100 steps?
*/
func (d *dumboOctopus) flashCount(input []string) int {
	grid := d.parseInput(input)

	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += grid.step()
	}
	return flashes
}

func (d *dumboOctopus) syncStep(input []string) int {
	grid := d.parseInput(input)
	totalOctopuses := grid.width*grid.length
	steps := 0
	for {
		steps++
		if flashes := grid.step(); flashes == totalOctopuses {
			break 
		}
	}
	return steps
}

type flashGrid struct {
	width  int
	length int
	octos  []int
}

func (f *flashGrid) step() int {
	flashed := make(map[int]struct{})
	var queue []int
	for i := 0; i < len(f.octos); i++ {
		f.octos[i]++
		if f.octos[i] > 9 {
			flashed[i] = struct{}{}
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		flash := queue[len(queue)-1]
		queue = queue[:len(queue)-1] //remove element from queue
		f.octos[flash] = 0
		for _, neighbor := range f.neighbors(flash) {
			if _, found := flashed[neighbor]; found {
				continue
			}
			f.octos[neighbor]++
			if f.octos[neighbor] > 9 {
				flashed[neighbor] = struct{}{}
				queue = append(queue, neighbor)
			}
		}
	}
	return len(flashed)
}

//changes from x, y to neighbors
var neighborDeltas = []point{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func (f *flashGrid) neighbors(i int) []int {
	var neighborPositions []int
	x, y := i%f.width, i/f.width
	for _, delta := range neighborDeltas {
		dx, dy := x+delta.x, y+delta.y
		if dx < 0 || dx >= f.width || dy < 0 || dy >= f.length {
			continue
		}

		neighborPositions = append(neighborPositions, dy*f.width+dx)
	}
	return neighborPositions
}

func (d *dumboOctopus) parseInput(input []string) *flashGrid {
	var f flashGrid
	f.length = len(input)
	f.width = len(input[0])
	for _, line := range input {
		for _, char := range line {
			f.octos = append(f.octos, int(char)-48)
		}
	}

	return &f
}
