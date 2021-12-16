package day11

import (
	"github.com/davidparks11/advent2021/internal/coordinate"
)

type FlashGrid struct {
	Width  int
	Length int
	octos  []int
}

func (f *FlashGrid) Step() int {
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
var neighborDeltas = []coordinate.Point{
	{-1, -1}, {0, -1}, {1, -1},
	{-1, 0}, {1, 0},
	{-1, 1}, {0, 1}, {1, 1},
}

func (f *FlashGrid) neighbors(i int) []int {
	var neighborPositions []int
	x, y := i%f.Width, i/f.Width
	for _, delta := range neighborDeltas {
		dx, dy := x+delta.X, y+delta.Y
		if dx < 0 || dx >= f.Width || dy < 0 || dy >= f.Length {
			continue
		}

		neighborPositions = append(neighborPositions, dy*f.Width+dx)
	}
	return neighborPositions
}

func ParseInput(input []string) *FlashGrid {
	var f FlashGrid
	f.Length = len(input)
	f.Width = len(input[0])
	for _, line := range input {
		for _, char := range line {
			f.octos = append(f.octos, int(char)-48)
		}
	}

	return &f
}

