package advent

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Problem interface {
	Solve()
}

type dailyProblem struct {
	day  int
	name string
}

//WriteResult takes result as a string and writes/overwrites the content to a result.txt file
func (d *dailyProblem) WriteResult(results []string) {
	fileName := fmt.Sprintf("resources/results/result%d.txt", d.day)
	resultFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		fmt.Sprintf(err.Error())
		return
	}

	for i, result := range results {
		fmt.Printf("Result for Day %d, the %s Problem, Part %d, : %v\n", d.day, d.name, i+1, result)
		resultFile.WriteString(result + "\n")

	}

	if err := resultFile.Close(); err != nil {
		fmt.Print(err.Error())
		return
	}
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
