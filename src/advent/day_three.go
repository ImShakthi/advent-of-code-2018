package advent

import (
	"fmt"
	"regexp"
	"strconv"
)

const size = 1000

// Matrix is a two dimensional array
type Matrix [size][size]int

var claimArr Matrix
var overLapArr [1342]int

//Claim struct
type Claim struct {
	claimID int
	yCoord  int
	xCoord  int
	cols    int
	rows    int
}

var pattern = "#(?P<CalimID>[0-9]*)\\s@\\s(?P<X>[0-9]*),(?P<Y>[0-9]*):\\s(?P<ROWS>[0-9]*)x(?P<COLS>[0-9]*)"
var claimMetaData = regexp.MustCompile(pattern)

// ClaimMatrix is function
func ClaimMatrix() {
	fmt.Println("Claim Matrix")
	data := ReadDataFromFile("day_three.txt")
	noL := len(data)
	counter := 0
	for index := 0; index < noL; index++ {
		claim := parseInput(data[index])
		updateArr(claim)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if claimArr[i][j] == -1 {
				counter++
			}
		}
	}
	fmt.Println("counter =", counter)

	//fmt.Println(claimArr)
	for i := 1; i < len(overLapArr); i++ {
		if overLapArr[i] == 0 {
			fmt.Print(" ", i)
		}
	}
}

// 00 01 02
// 10 11 12
// 20 21 22

func updateArr(claim Claim) {
	//fmt.Println(claim)
	for i := claim.xCoord; i < (claim.xCoord + claim.rows); i++ {
		for j := claim.yCoord; j < (claim.yCoord + claim.cols); j++ {
			if claimArr[i][j] == 0 {
				claimArr[i][j] = claim.claimID
			} else {
				overLapArr[claim.claimID] = 1
				if claimArr[i][j] != -1 {
					overLapArr[claimArr[i][j]] = 1
				}
				claimArr[i][j] = -1
			}
		}
	}
}

func parseInput(inputStr string) (claim Claim) {
	matches := claimMetaData.FindStringSubmatch(inputStr)
	claim = Claim{convInt(matches[1]),
		convInt(matches[2]),
		convInt(matches[3]),
		convInt(matches[4]),
		convInt(matches[5]),
	}
	return
}

func convInt(data string) (number int) {
	number, _ = strconv.Atoi(data)
	return
}
