package advent

import (
	. "github.com/davidparks11/advent2021/internal/advent/day21"
	"github.com/davidparks11/advent2021/internal/math"
)

type diracDice struct {
    dailyProblem
}

func NewDiracDice() Problem {
    return &diracDice{
        dailyProblem{
            day: 21,
        },
    }
}

func (d *diracDice) Solve() interface{} {
    input := d.GetInputLines()
    var results []int
    results = append(results, d.part1(input))
    results = append(results, d.part2(input))
    return results
}

func (d *diracDice) part1(input []string) int {
    p1, p2 := ParseInput(input)
    regularDie := NewDie(100)
    var loser *Player
    for {
        p1.Move(regularDie.RollMany(3))
        if p1.Score >= 1000 {
            loser = p2
            break
        }
        p2.Move(regularDie.RollMany(3))
        if p2.Score >= 1000 {
            loser = p1
            break
        }
    }

    return loser.Score * regularDie.TimesRolled
}

func (d *diracDice) part2(input []string) int {
    p1, p2 := ParseInput(input)
    return math.Max(QuantumDiceGame(0, 0, p1.Pos, p2.Pos)) 
}
