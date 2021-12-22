package day18

import (
	"math"
)

type Node struct {
	value  *int
	parent *Node
	right  *Node
	left   *Node
}

type doneTraversingFunc func(s *Node, d int) bool

func (n *Node) copy() *Node {
	root := &Node{value: n.value}
	if n.left != nil {
		root.left = n.left.copy()
		root.left.parent = root
	}
	if n.right != nil {	
		root.right = n.right.copy()
		root.right.parent = root
	}

	return root
}

func AddNode(n1, n2 *Node) *Node {
	newRoot := &Node{left: n1, right: n2}
	n1.parent = newRoot
	n2.parent = newRoot
	newRoot = reduce(newRoot)
	return newRoot
}

func reduce(n *Node) *Node {
	var reduced bool 
	for !reduced {
		reduced = true
		toExplode := nextExplode(n)
		for toExplode != nil {
			reduced = false
			explode(toExplode)
			toExplode = nextExplode(n)
		}

		toSplit := nextSplit(n)
		if toSplit != nil {
			reduced = false
			split(toSplit)
		}
	}

	return n
}

func explode(n *Node) {
	nextLeft := getNextLeft(n)
	if nextLeft != nil {
		if nextLeft.value == nil {
			nextLeft = rightChild(nextLeft)
		}
		leftValue := addValues(n.left, nextLeft)
		nextLeft.value = leftValue
	}

	nextRight := getNextRight(n)
	if nextRight != nil {
		if nextRight.value == nil {
			nextRight = leftChild(nextRight)
		}
		rightValue := addValues(n.right, nextRight)
		nextRight.value = rightValue
	}

	value := 0
	newNode := &Node{value: &value, parent: n.parent}
	if n == n.parent.right {
		newNode.parent.right = newNode
	} else if n == n.parent.left {
		newNode.parent.left = newNode
	}

	removeNode(n.left)
	removeNode(n.right)
	removeNode(n)
}

func removeNode(n *Node) {
	if n == nil {
		return
	}
	if n == n.parent.left {
		n.parent.left = nil
	}
	if n == n.parent.right {
		n.parent.right = nil
	}
	n.parent = nil
	n.left = nil
	n.right = nil
}

func split(n *Node) {
	v1 := *n.value / 2
	v2 := *n.value - v1
	n.value = nil

	left := &Node{value: &v1, parent: n}
	right := &Node{value: &v2, parent: n}

	n.left = left
	n.right = right
}

func leftChild(n *Node) *Node {
	for n.left != nil {
		n = n.left
	}
	return n
}

func rightChild(n *Node) *Node {
	for n.right != nil {
		n = n.right
	}
	return n
}

func getNextLeft(current *Node) *Node {
	n := current
	for current != nil {
		current = current.parent
		if current != nil && current.left != nil && current.left != n {
			return current.left
		}
		n = current
	}
	return nil
}

func getNextRight(current *Node) *Node {
	n := current
	for current != nil {
		current = current.parent
		if current != nil && current.right != nil && current.right != n {
			return current.right
		}
		n = current
	}
	return nil
}

func addValues(n1, n2 *Node) *int {
	if n1 == nil || n2 == nil || n1.value == nil || n2.value == nil {
		return nil
	}
	value := *n1.value + *n2.value
	return &value
}

func nextSplit(n *Node) (toSplit *Node) {
	var checkSplit doneTraversingFunc = func(c *Node, d int) bool {
		if c.value != nil && *c.value > 9 && toSplit == nil {
			toSplit = c
			return true
		}
		return false
	}

	traverse(n, checkSplit, 0)
	return
}

func nextExplode(n *Node) (toExplode *Node) {
	var checkExplode doneTraversingFunc = func(n *Node, d int) bool {
		if d >= 4 && n.left != nil && n.left.value != nil && n.right != nil && n.right.value != nil {
			toExplode = n
			return true
		}
		return false
	}

	traverse(n, checkExplode, 0)
	return 
}

func traverse(n *Node, finished doneTraversingFunc, depth int) bool {
	if finished(n, depth) {
		return true
	}
	if n.left != nil {
		if traverse(n.left, finished, depth+1) {
			return true
		}
	}
	if n.right != nil {
		if traverse(n.right, finished, depth+1) {
			return true
		}
	}
	return false
}

func MaxMagnitude(list []*Node) int {
	var max = math.MinInt

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if i == j {
				continue
			}
			n1, n2 := list[i].copy(), list[j].copy()
			newRoot := AddNode(n1, n2)
			m := CalcMagnitude(newRoot)
			if m > max {
				max = m
			}
		}
	}

	return max
}

func CalcMagnitude(n *Node) int {
	if n == nil {
		return 0
	}
	if n.value != nil {
		return *n.value
	}
	return 3*CalcMagnitude(n.left) + 2*CalcMagnitude(n.right)
}

func ParseInput(input []string) []*Node {

	parseNode := func(str string) *Node {
		root := &Node{}
		current := root
		for _, char := range str {
			switch char {
			case '[':
				n := &Node{parent: current}
				current.left = n
				current = n
			case ']':
				current = current.parent
			case ',':
				n := &Node{parent: current}
				current.right = n
				current = n
			default:
				value := int(char - '0')
				current.value = &value
				current = current.parent
			}
		}
		return root
	}

	var nodes []*Node
	for _, line := range input {
		n := parseNode(line)
		nodes = append(nodes, n)
	}
	return nodes
}
