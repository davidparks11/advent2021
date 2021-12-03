package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Problem interface {
	Solve() []string
	Day() int
}

type dailyProblem struct {
	day int
}

func (d *dailyProblem) Day() int {
	return d.day
}

//GetInputLines reads an input.txt file and returns its contents separated by lines as a string array
func (d *dailyProblem) GetInputLines() []string {
	fileName := fmt.Sprintf("resources/inputs/input%d.txt", d.day)
	inputFile, err := os.Open(fileName)
	if err != nil {
		fmt.Print(err.Error())
		return nil
	}

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	inputFile.Close()

	return lines
}

//IntsFromStrings takes a string array and returns array of those strings converted to ints
func IntsFromStrings(inputLines []string) []int {
	input := make([]int, len(inputLines))
	for i, line := range inputLines {
		intValue, err := strconv.Atoi(line)
		if err != nil {
			return []int{}
		}
		input[i] = intValue
	}
	return input
}
