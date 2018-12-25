package advent

import (
	"fmt"
	"regexp"
)

// MarbleMania is Elf play area
func MarbleMania() {
	fmt.Println("--- Marble Mania ---")

	input := ""
	// input = "10 players; last marble is worth 1618 points"
	input = "13 players; last marble is worth 7999 points"
	// input = "17 players; last marble is worth 1104 points"
	// input = "21 players; last marble is worth 6111 points"
	input = "30 players; last marble is worth 5807 points"
	input = "430 players; last marble is worth 71588 points"
	// input = "9 players; last marble is worth 25 points"

	var pattern = "(?P<players>[0-9]*) players; last marble is worth (?P<marbles>[0-9]*) points"
	matches := regexp.MustCompile(pattern).FindStringSubmatch(input)

	nosPlayers := ConvInt(matches[1])
	nosMarbles := ConvInt(matches[2])
	// fmt.Println("# of players =", nosPlayers, ", # of marbles =", nosMarbles)

	playerScoreMap := make(map[int]int)
	marbleList := []int{0}

	placeIndex := 1
	player := -1

	for i := 1; i <= nosMarbles; i++ {
		player = (player + 1) % nosPlayers

		if i%23 == 0 {
			playerScoreMap[player] += i

			placeIndex = (placeIndex - 7 - 2) % len(marbleList)
			if placeIndex < 0 {
				placeIndex = -placeIndex
			}
			// fmt.Println("placeIndex =", placeIndex)
			playerScoreMap[player] += marbleList[placeIndex]
			marbleList = RemoveAtIndex(marbleList, placeIndex)
		} else {
			marbleList = AppendAtIndex(marbleList, placeIndex, i)
		}

		placeIndex += 2
		if len(marbleList) == placeIndex && ((len(marbleList) % 2) == 0) {
			//fmt.Println("placeIndex =", placeIndex)
			placeIndex = 0
		} else if len(marbleList) < placeIndex {
			placeIndex = 1
			//placeIndex % (len(marbleList))
		}

		// fmt.Println("Player #", player, "-", marbleList)
	}

	// fmt.Println(playerScoreMap)
	max := 0
	for _, score := range playerScoreMap {
		if max < score {
			max = score
		}
	}
	fmt.Println("max score =", max)
}
