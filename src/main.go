package main

import (
	"advent"
	"basics"
	"fmt"
	"rectangle"
	//"snippets"
)

func main() {
	adventOfCode()
}

func adventOfCode() {
	day := 5
	switch day {
	case 1:
		advent.CalcFrequency()
	case 2:
		advent.CalcThreshold()
	case 3:
		advent.ClaimMatrix()
	case 4:
		advent.ReposeRecord()
	case 5:
		advent.AlchemicalReduction()
	default:
	}
}

