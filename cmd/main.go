package main

import (
	"flag"

	"github.com/davidparks11/advent2021/internal/advent"
)

//TODO Add usage


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

	problems := advent.NewProblemSet()
	problems.Solve(printToConsole, day)

}

