package day4

import (
	"strconv"
	"strings"
)

var winMasks = []uint32{
	0b1111100000000000000000000, //first row
	0b0000011111000000000000000, //second row
	0b0000000000111110000000000, //third row
	0b0000000000000001111100000, //fourth row
	0b0000000000000000000011111, //fifth row
	0b1000010000100001000010000, //first column
	0b0100001000010000100001000, //second column
	0b0010000100001000010000100, //third column
	0b0001000010000100001000010, //fourth column
	0b0000100001000010000100001, //fifth column
}

type Board struct {
	places  [25]int
	marks   uint32
	lastNum int
	won     bool
}

func (b *Board) Mark(num int) bool {
	if b.won {
		return false //prevent returning a board more than once as a winner
	}

	for i, place := range b.places {
		if place == num {
			b.lastNum = num
			b.marks |= (1 << i)
			for _, mask := range winMasks {
				if b.marks&mask == mask {
					b.won = true
					return true
				}
			}
		}
	}

	return false
}

func (b *Board) Score() int {
	unused := 0
	for i, num := range b.places {
		if b.marks&(1<<i) == 0 {
			unused += num
		}
	}

	return unused * b.lastNum
}


func MakeBoards(input []string) []*Board {
	var boards []*Board
	for i := 0; i < len(input); i += 6 {
		boards = append(boards, makeBoard(input[i:i+5]))
	}
	return boards
}

func makeBoard(input []string) *Board {
	var b Board
	for i, line := range input {
		line = strings.TrimSpace(line)
		line = strings.ReplaceAll(line, "  ", " ")
		numStrings := strings.Split(line, " ")
		for j, numString := range numStrings {
			num, err := strconv.ParseInt(numString, 10, 32)
			if err != nil {
				panic(err.Error())
			}
			b.places[i*5+j] = int(num)
		}
	}
	return &b
}


