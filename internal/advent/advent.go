package advent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Problem interface {
	Solve() []int
	Day() int
}

type problemSet map[int]Problem

func NewProblemSet() *problemSet {
	problems := []Problem{
		NewSonorSweep(),
		NewDive(),
		NewBinaryDiagnostic(),
		NewGiantSquid(),
		NewHydrothermalVenture(),
		NewLanternFish(),
		NewTheTreacheryOfWhales(),
		NewSmokeBasin(),
		NewSyntaxScoring(),
		NewDumboOctopus(),
	}

	p := make(problemSet)
	for _, problem := range problems {
		p[problem.Day()] = problem
	}

	return &p
}

func (p *problemSet) Get(day int) Problem {
	problem, found := (*p)[day]
	if !found {
		log.Fatal(fmt.Sprintf("problem not found in problem set: %d", day))
	}
	return problem
}

const Christmas = 25

func (p *problemSet) Solve(writeToConsole bool, day int) {
	if day != 0 {
		if writeToConsole {
			log.Printf("Result for Day %d: %v\n", day, p.Get(day).Solve())
		} else {
			p.WriteResultFile(day)
		}
	} else {
		for day := 1; day <= Christmas; day++ {
			if _, found := (*p)[day]; found {
				if writeToConsole {
					log.Printf("Result for Day %d: %v\n", day, p.Get(day).Solve())
				} else {
					p.WriteResultFile(day)
				}
			}
		}
	}
}

//WriteResult takes result as a string and writes/overwrites the content to a result.txt file
func (p *problemSet) WriteResultFile(day int) {
	problem := p.Get(day)

	fileName := fmt.Sprintf("resources/results/result%d.txt", day)
	resultFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = resultFile.WriteString(fmt.Sprint(problem.Solve())); err != nil {
		log.Fatal(err)
	}

	if err = resultFile.Close(); err != nil {
		log.Fatal(err)
	}
}

type dailyProblem struct {
	day int
}

func (d *dailyProblem) Day() int {
	return d.day
}

//GetInputLines reads an input.txt file and returns its contents separated by lines as a string array
func (d *dailyProblem) GetInputLines() []string {
	if d.day == 0 {
		log.Fatal("error getting input lines with no set day")
	}
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

func CommaSplitInts(input string) []int {
	var nums []int
	for _, numString := range strings.Split(input, ",") {
		num, err := strconv.ParseInt(numString, 10, 32)
		if err != nil {
			log.Fatal(err.Error())
		}
		nums = append(nums, int(num))
	}
	return nums
}
