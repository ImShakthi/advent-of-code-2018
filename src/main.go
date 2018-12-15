package main

import (
	"advent"
	"fmt"
	//"snippets"
)

func main() {
	fmt.Println("------------------------------------------------------------------------")
	adventOfCode()
	fmt.Println("------------------------------------------------------------------------")
}

func adventOfCode() {
	day := 5
	fmt.Println("--- Advent of code --- DAY #", day)
	switch day {
	case 1:
		advent.ChronalCalibration()
	case 2:
		advent.InventoryManagementSystem()
	case 3:
		advent.NoMatterHowYouSliceIt()
	case 4:
		advent.ReposeRecord()
	case 5:
		advent.AlchemicalReduction()
	case 6:
		advent.ChronalCoordinates()
	default:
	}
}
