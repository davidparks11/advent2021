package day12

type Graph struct {
	Edges map[Node][]Node
	Start Node
	End Node
}

type Node string

