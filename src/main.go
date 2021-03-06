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
	day := 9
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
	case 7:
		advent.TheSumofItsParts()
	case 8:
		advent.MemoryManeuver()
	case 9:
		advent.MarbleMania()
	default:
	}
}
