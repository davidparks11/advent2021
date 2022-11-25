package day24

import (
	"strconv"

	"github.com/davidparks11/advent2021/internal/math"
)

type instruction struct {
	div  int
	xInc int
	yInc int
}

const BLOCK_SIZE = 18

func FindModelNumber(instructions []string, isMax bool) int {
	_, ans := recurse(parseInput(instructions), isMax)
	return digitsToInt(ans)
}


func recurse(instructions []instruction, isMax bool) ([]instruction, []int) {
	if len(instructions) == 0 || instructions[0].div == 26 {
		return instructions, []int{}
	}

	left, instructions := pop(instructions)
	instructions, mid := recurse(instructions, isMax)
	right, instructions := pop(instructions)
	var leftOutput int
	if isMax {
		leftOutput = math.Min(9, 9-left.yInc-right.xInc)
	} else {
		leftOutput = math.Max(1, 1-left.yInc-right.xInc)
	}
	rightOutput := leftOutput + left.yInc + right.xInc
	instructions, tail := recurse(instructions, isMax)

	return instructions, appendArrays([]int{leftOutput}, mid, []int{rightOutput}, tail)
}

func pop[T any](arr []T) (T, []T) {
	var v T
	if len(arr) == 0 {
		return v, []T{}
	}

	v = arr[0]
	arr = arr[1:]
	return v, arr
}

func appendArrays[T any](arrays ...[]T) []T {
	var length int
	for _, a := range arrays {
		length += len(a)
	}

	arr := make([]T, length)
	i := 0
	for _, a := range arrays {
		for _, v := range a {
			arr[i] = v
			i++
		}
	}

	return arr
}

func parseInput(input []string) []instruction {
	blocks := len(input) / BLOCK_SIZE
	instructions := make([]instruction, blocks)

	for i := 0; i < blocks; i++ {
		instructions[i] = newInstruction(input[i*BLOCK_SIZE+4], input[i*BLOCK_SIZE+5], input[i*BLOCK_SIZE+15])
	}

	return instructions
}

func newInstruction(divInput, xIncInput, yIncInput string) instruction {
	i := instruction{}

	if len(divInput) == 7 {
		i.div = 1
	} else {
		i.div = 26
	}

	if v, err := strconv.Atoi(xIncInput[6:]); err != nil {
		panic("ALU: bad int data - " + err.Error())
	} else {
		i.xInc = v
	}

	if v, err := strconv.Atoi(yIncInput[6:]); err != nil {
		panic("ALU: bad int data - " + err.Error())
	} else {
		i.yInc = v
	}

	return i
}

func digitsToInt(digits []int) (v int) {
	for _, d := range digits {
		v *= 10
		v += d
	}
	return
}
