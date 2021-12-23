package day20

import (
	"github.com/davidparks11/advent2021/internal/coordinate"
)

type mapData struct {
	enhancementAlg         []bool
	img                    map[coordinate.Point]bool
	defaultLit             bool
	yMin, yMax, xMin, xMax int
}

func (m *mapData) String() string {
	str := ""
	renderBool := func(b bool) string {
		if b {
			return "#"
		}
		return "."
	}

	for _, b := range m.enhancementAlg {
		str += renderBool(b)
	}

	str += "\n\n"

	for row := 0; row < m.yMax; row++ {
		for col := 0; col < m.xMax; col++ {
			str += renderBool(m.img[coordinate.Point{Y: row, X: col}])
		}
		str += "\n"
	}

	return str
}

func (m *mapData) LitCount() (count int) {
	for _, lit := range m.img {
		if lit {
			count++
		}
	}
	return
}

func (m *mapData) Enhance() {
	m.expand()
	// fmt.Println(m)
	newImg := make(map[coordinate.Point]bool)
	for row := m.yMin; row < m.yMax; row++ {
		for col := m.xMin; col < m.xMax; col++ {
			pixel := coordinate.Point{Y: row, X: col}
				newImg[pixel] = m.enhancementValue(row, col)
		}
	}

	if m.defaultLit {
		m.defaultLit = m.enhancementAlg[511]
	} else {
		m.defaultLit = m.enhancementAlg[0]
	}

	m.img = newImg
}

var neighbors = []coordinate.Point{
	{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1},
	{X: -1, Y: 0}, {X: 0, Y: 0}, {X: 1, Y: 0},
	{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1},
}

func (m *mapData) enhancementValue(row, col int) bool {
	index := 0
	for i, delta := range neighbors {
		if lit, found := m.img[coordinate.Point{Y: row + delta.Y, X: col + delta.X}]; lit || (!found && m.defaultLit) {
			index |= 1 << (8 - i)
		}
	}
	value := m.enhancementAlg[index]
	// fmt.Println("row: ", row, "\tcol: ", col, "\tindex: ", index, "\tvalue: ", value)
	return value
}

func (m *mapData) expand() {
	m.yMin--
	m.yMax++
	m.xMin--
	m.xMax++
}

const inputImgStart = 2

func ParseInput(input []string) *mapData {
	var data mapData
	data.enhancementAlg = make([]bool, len(input[0]))
	for i, char := range input[0] {
		if char == '#' {
			data.enhancementAlg[i] = true
		}
	}

	// if data.enhancementAlg[0] {
	// 	data.defaultLit = true
	// }

	data.xMin = 0
	data.yMin = 0
	data.yMax = len(input) - inputImgStart
	data.xMax = len(input[inputImgStart])

	data.img = make(map[coordinate.Point]bool)

	for y, line := range input[inputImgStart:] {
		for x, char := range line {
			if char == '#' {
				data.img[coordinate.Point{X: x, Y: y}] = true
			}
		}
	}

	return &data
}
