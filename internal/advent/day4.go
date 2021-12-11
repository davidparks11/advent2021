package advent

import (
	"strconv"
	"strings"
)

type giantSquid struct {
	dailyProblem
}

func NewGiantSquid(day int) Problem {
	return &giantSquid{
		dailyProblem{
			day: day,
		},
	}
}

func (g *giantSquid) Solve() []string {
	input := g.GetInputLines()
	var results []string
	results = append(results, strconv.Itoa(g.winningBoardScore(input)))
	return results
}

/*
You're already almost 1.5km (almost a mile) below the surface of the ocean, already so deep that you can't see any sunlight. What you can see, however, is a giant squid that has attached itself to the outside of your submarine.

Maybe it wants to play bingo?

Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the chosen number is marked on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in any row or any column of a board are marked, that board wins. (Diagonals don't count.)

The submarine has a bingo subsystem to help passengers (currently, you and the giant squid) pass the time. It automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For example:

7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
Finally, 24 is drawn:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7
At this point, the third board wins because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: 14 21 17 24 4).

The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.

To guarantee victory against the giant squid, figure out which board will win first. What will your final score be if you choose that board?
*/
func (g *giantSquid) winningBoardScore(input []string) int {
	chosenNums := g.parseChosenNumbers(input[0])
	boards := g.makeBoards(input[2:])
	for i, num := range chosenNums {
		for _, b := range boards {
			if winner := b.mark(num); i > 4 && winner {
				return b.score(num)
			}
		}
	}
	return 0
}

type board struct {
	places [25]int
	marks uint32
}

func (b *board) mark(num int) bool {
	for i, place := range b.places {
		if place == num {
			b.marks |= (1 << i)
			for _, mask := range winMasks {
				if b.marks & mask == mask {
					return true
				}
			}
		}
	}

	return false
}

func (b *board) score(lastNum int) int {
	unused := 0
	for i, num := range b.places {
		if b.marks & (1 << i) == 0 {
			unused += num
		}
	}

	return unused * lastNum
}

var winMasks = []uint32 {
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

func (g *giantSquid) parseChosenNumbers(input string) []int {
	var nums []int
	numStrings := strings.Split(input, ",")
	for _, numString := range numStrings {
		num, err := strconv.ParseInt(numString, 10, 32)
		if err != nil {
			panic(err.Error())
		}
		nums = append(nums, int(num))
	}
	return nums
}

func (g *giantSquid) makeBoards(input []string) []*board {
	var boards []*board
	for i := 0; i < len(input); i += 6 {
		boards = append(boards, g.makeBoard(input[i:i+5]))
	}
	return boards
}

func (g *giantSquid) makeBoard(input []string) *board {
	var b board
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
