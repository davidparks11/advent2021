package day21

import (
	"strconv"
	"strings"
)

//a mapping of roll counts to the number of ways that a roll can be reached.
//For example, there are 3 ways to roll an '8' -> (2, 3, 3), (3, 2, 3), (3, 3, 2)
var combinationsByTotals = map[int]int{3:1, 4:3, 5:6, 6:7, 7:6, 8:3, 9:1}

func QuantumDiceGame(p1Score, p2Score, p1Pos, p2Pos int) (int, int) {
	//check win conditions
    if p1Score >= 21 {
        return 1, 0
    }
	
    if p2Score >= 21 {
        return 0, 1
    }

	//track player 1 and 2 wins
    p1totalWins, p2TotalWins := 0, 0

	//loop through all possible rolls, calling quantumDiceGame() recursively but replacing p1 with p2
    for totalValue, combo := range combinationsByTotals {
		nextPos := (p1Pos + totalValue - 1)%10 + 1
        p2Wins, p1Wins := QuantumDiceGame(p2Score, p1Score + nextPos, p2Pos, nextPos) //switch inputs - so switch expected universes
		//multiple wins by combinations for a roll
        p1totalWins += combo * p1Wins
        p2TotalWins += combo * p2Wins
    }
    return p1totalWins, p2TotalWins
}

func ParseInput(input []string) (p1 *Player, p2 *Player) {
    pos1, err := strconv.Atoi(strings.Split(input[0], ": ")[1])
    if err != nil {
        panic(err)
    }

    pos2, err := strconv.Atoi(strings.Split(input[1], ": ")[1])
    if err != nil {
        panic(err)
    }

    p1 = &Player{
        Pos: pos1,
    }
    p2 = &Player{
        Pos: pos2,
    }
    return
}

type Die struct {
    sides int
    numberRolled int
    TimesRolled int
}

func NewDie(sides int) *Die {
    return &Die{
        sides: sides,
        numberRolled: 1,
    }
}

func (d *Die) RollMany(times int) int {
    total := 0
    for i := 0; i < times; i++ {
        total += d.roll()
    }
    return total
}

func (d *Die) roll() int {
    val := d.numberRolled
    d.numberRolled = (d.numberRolled)%d.sides + 1
    d.TimesRolled++
    return val
}

type Player struct {
    Pos   int
    Score int
}

func (p *Player) Move(spaces int) {
    p.Pos = (p.Pos + spaces - 1)%10 + 1
    p.Score += p.Pos
}