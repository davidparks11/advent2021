package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/davidparks11/advent2021/internal/advent"
)

//TODO Add usage

var problems = []advent.Problem{
	advent.NewSonorSweep(1),
	advent.NewDive(2),
	advent.NewBinaryDiagnostic(3),
	advent.NewSmokeBasin(9),
	advent.NewSyntaxScoring(10),
}

func main() {
	var printToConsole bool
	var day int

	inputPrintToConsole := flag.Bool("console", false, "prints to console instead of a result file") 
	inputDay := flag.Int("day", 0, "result for specific day 1-25") 
	
	flag.Parse()

	if inputPrintToConsole != nil {
		printToConsole = *inputPrintToConsole
	}

	if inputDay != nil && *inputDay >= 1 && *inputDay <= 25 {
		day = *inputDay
	}

	for _, p := range problems {
		problemDay := p.Day()
		if day == 0 || problemDay == day {
			result := p.Solve()
			if printToConsole {
				log.Printf("Result for Day %d: %v\n", problemDay, result)
			} else {
				WriteResult(result, problemDay)
			}
		} 
	}
}

//WriteResult takes result as a string and writes/overwrites the content to a result.txt file
func WriteResult(result []string, day int) {
	fileName := fmt.Sprintf("resources/results/result%d.txt", day)
	resultFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = resultFile.WriteString(strings.Join(result, "\n")); err != nil {
		log.Fatal(err)
	}
	
	if err = resultFile.Close(); err != nil {
		log.Fatal(err)
	}
}

