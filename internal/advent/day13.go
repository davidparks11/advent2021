package advent

import (
	"strconv"
	"strings"
)

var _ Problem = &transparentOrigami{}

type transparentOrigami struct {
	dailyProblem
}

func NewTransparentOrigami() Problem {
	return &transparentOrigami{
		dailyProblem{
			day: 13,
		},
	}
}

func (d *transparentOrigami) Solve() interface{} {
	input := d.GetInputLines()
	var results []string
	results = append(results, strconv.FormatInt(int64(d.countPoints(input)), 10))
	results = append(results, d.renderText(input))

	return results
}

/*
You reach another volcanically active part of the cave. It would be nice if you could do some kind of thermal imaging so you could tell ahead of time which caves are too hot to safely enter.

Fortunately, the submarine seems to be equipped with a thermal camera! When you activate it, you are greeted with:

Congratulations on your purchase! To activate this infrared thermal imaging
camera system, please enter the code found on page 1 of the manual.
Apparently, the Elves have never used this feature. To your surprise, you manage to find the manual; as you go to open it, page 1 falls out. It's a large sheet of transparent paper! The transparent paper is marked with random dots and includes instructions on how to fold it up (your puzzle input). For example:

6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
The first section is a list of dots on the transparent paper. 0,0 represents the top-left coordinate. The first value, x, increases to the right. The second value, y, increases downward. So, the coordinate 3,0 is to the right of 0,0, and the coordinate 0,7 is below 0,0. The coordinates in this example form the following pattern, where # is a dot on the paper and . is an empty, unmarked position:

...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
...........
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
Then, there is a list of fold instructions. Each instruction indicates a line on the transparent paper and wants you to fold the paper up (for horizontal y=... lines) or left (for vertical x=... lines). In this example, the first fold instruction is fold along y=7, which designates the line formed by all of the positions where y is 7 (marked here with -):

...#..#..#.
....#......
...........
#..........
...#....#.#
...........
...........
-----------
...........
...........
.#....#.##.
....#......
......#...#
#..........
#.#........
Because this is a horizontal line, fold the bottom half up. Some of the dots might end up overlapping after the fold is complete, but dots will never appear exactly on a fold line. The result of doing this fold looks like this:

#.##..#..#.
#...#......
......#...#
#...#......
.#.#..#.###
...........
...........
Now, only 17 dots are visible.

Notice, for example, the two dots in the bottom left corner before the transparent paper is folded; after the fold is complete, those dots appear in the top left corner (at 0,0 and 0,1). Because the paper is transparent, the dot just below them in the result (at 0,3) remains visible, as it can be seen through the transparent paper.

Also notice that some dots can end up overlapping; in this case, the dots merge together and become a single dot.

The second fold instruction is fold along x=5, which indicates this line:

#.##.|#..#.
#...#|.....
.....|#...#
#...#|.....
.#.#.|#.###
.....|.....
.....|.....
Because this is a vertical line, fold left:

#####
#...#
#...#
#...#
#####
.....
.....
The instructions made a square!

The transparent paper is pretty big, so for now, focus on just completing the first fold. After the first fold in the example above, 17 dots are visible - dots that end up overlapping after the fold is completed count as a single dot.

How many dots are visible after completing just the first fold instruction on your transparent paper?
*/
func (t *transparentOrigami) countPoints(input []string) int {
	points, folds := t.parseInput(input)
	fold := folds[0]
	for p, _ := range points {
		if fold.x != 0 && fold.x < p.x {
			points[point{x: fold.x - (p.x - fold.x), y: p.y}] = struct{}{}
			delete(points, p)
		}
		if fold.y != 0 && fold.y < p.y {
			points[point{x: p.x, y: fold.y - (p.y - fold.y)}] = struct{}{}
			delete(points, p)
		}
	}

	return len(points)
}

/*
Finish folding the transparent paper according to the instructions. The manual says the code is always eight capital letters.

What code do you use to activate the infrared thermal imaging camera system?
*/
func (t *transparentOrigami) renderText(input []string) string {
	points, folds := t.parseInput(input)
	for _, f := range folds {
		for p, _ := range points {
			if f.x != 0 && f.x < p.x {
				points[point{x: f.x - (p.x - f.x), y: p.y}] = struct{}{}
				delete(points, p)
			}
			if f.y != 0 && f.y < p.y {
				points[point{x: p.x, y: f.y - (p.y - f.y)}] = struct{}{}
				delete(points, p)
			}
		}
	}

	return t.renderLetters(points)
}

func (t *transparentOrigami) renderLetters(points map[point]struct{}) string {
	var width int
	var length int
	for p, _ := range points {
		if p.x+1 > width {
			width = p.x + 1
		}
		if p.y+1 > length {
			length = p.y + 1
		}
	}

	output := ""

	for i := 0; i <= width*length; i++ {
		if i != 0 && i%width == 0 {
			output += "\n"
		}

		if _, found := points[point{x: i % width, y: i / width}]; found {
			output += "█"
		} else {
			output += " "
		}

	}

	return output
}

func (t *transparentOrigami) parseInput(input []string) (points map[point]struct{}, folds []point) {
	points = make(map[point]struct{})

	i := 0
	for ; input[i] != ""; i++ {
		pointStrings := strings.Split(input[i], ",")
		x, err := strconv.ParseInt(pointStrings[0], 10, 32)
		if err != nil {
			panic(err.Error())
		}
		y, err := strconv.ParseInt(pointStrings[1], 10, 32)
		if err != nil {
			panic(err.Error())
		}
		points[point{x: int(x), y: int(y)}] = struct{}{}
	}

	for i++; i < len(input); i++ {
		lineStrings := strings.Split(input[i][11:], "=")
		var p point
		val, err := strconv.ParseInt(lineStrings[1], 10, 32)
		if err != nil {
			panic(err.Error())
		}

		if lineStrings[0] == "x" {
			p = point{x: int(val)}
		} else {
			p = point{y: int(val)}
		}

		folds = append(folds, p)
	}

	return
}
