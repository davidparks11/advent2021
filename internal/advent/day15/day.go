package day15

import (
	"container/heap"
	"github.com/davidparks11/advent2021/internal/coordinate"
)

type Grid [][]int

func (g Grid) InBounds(p coordinate.Point) bool {
	return p.Y >= 0 && p.X >= 0 && p.Y < len(g) && p.X < len(g[0])
}

//Expand returns a new grid of size x * y * n^2 with copies of itself in the x and y direction n times. All values in
//the new grid are incremented but are between 1-9
func (g Grid) Expand(n int) Grid {
	expanded := make(Grid, len(g)*n)

	for y := 0; y < g.Length()*n; y++ {
		for x := 0; x < g.Width()*n; x++ {
			origVal := g[y%g.Length()][x%g.Width()]
			newVal := (origVal+x/g.Width()+y/g.Length()-1)%9 + 1 //limit values to 1-9
			expanded[y] = append(expanded[y], newVal)
		}
	}

	return expanded
}

func (g Grid) Width() int  { return len(g[0]) }
func (g Grid) Length() int { return len(g) }

type Node struct {
	coordinate.Point
	Cost int
}

var _ heap.Interface = &NodeHeap{}

// An NodeHeap is a min-heap of nodes.
type NodeHeap []*Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Cost < h[j].Cost }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
