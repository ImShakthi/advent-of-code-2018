package advent

import (
	"fmt"
	"regexp"
)

type marble struct {
	value int
	next  *marble
	prev  *marble
}

// MarbleMania is Elf play area
func MarbleMania() {
	fmt.Println("--- Marble Mania ---")

	input := ""
	// input = "10 players; last marble is worth 1618 points"
	// input = "13 players; last marble is worth 7999 points"
	// input = "17 players; last marble is worth 1104 points"
	// input = "21 players; last marble is worth 6111 points"
	// input = "30 players; last marble is worth 5807 points"
	input = "430 players; last marble is worth 7158800 points" //3412522480
	// input = "9 players; last marble is worth 25 points"

	var pattern = "(?P<players>[0-9]*) players; last marble is worth (?P<marbles>[0-9]*) points"
	matches := regexp.MustCompile(pattern).FindStringSubmatch(input)

	nosPlayers := ConvInt(matches[1])
	nosMarbles := ConvInt(matches[2])
	// fmt.Println("# of players =", nosPlayers, ", # of marbles =", nosMarbles)

	// playerScoreMap := getElfScores(nosPlayers, nosMarbles)
	playerScoreMap := getElfScoresCLL(nosPlayers, nosMarbles)

	//fmt.Println(playerScoreMap)

	fmt.Println("max score =", getMaxScore(playerScoreMap))
}

func getElfScoresCLL(nosPlayers, nosMarbles int) map[int]int {
	fmt.Println("# of players =", nosPlayers, ", # of marbles =", nosMarbles)

	playerScoreMap := make(map[int]int)
	player := 0

	header := &marble{value: 0}
	nextElem := &marble{value: 1}

	header.next = nextElem
	header.prev = nextElem

	nextElem.next = header
	nextElem.prev = header

	// updating header

	for currentMarble := 2; currentMarble <= nosMarbles; currentMarble++ {
		player = (player + 1) % nosPlayers

		if currentMarble%23 == 0 {
			playerScoreMap[player] += currentMarble
			for i := 0; i < 7; i++ {
				header = header.prev
			}
			playerScoreMap[player] += header.value

			removeNode := header
			removeNode.next.prev = removeNode.prev
			removeNode.prev.next = removeNode.next

			header = removeNode.next
		} else {
			newNode := &marble{value: currentMarble,
				next: header.next.next, // to jump two nodes
				prev: header.next}
			newNode.prev.next = newNode
			newNode.next.prev = newNode
			header = newNode
		}
	}

	return playerScoreMap
}

func print(header *marble) {
	temp := header.next
	for i := 0; temp != header; i++ {
		fmt.Print(temp.value, ", ")
		temp = temp.next
	}
	fmt.Print(temp.value, "\n")
}
func getMaxScore(playerScoreMap map[int]int) int {
	max := 0
	for _, score := range playerScoreMap {
		if max < score {
			max = score
		}
	}
	return max
}
func getElfScores(nosPlayers, nosMarbles int) map[int]int {
	playerScoreMap := make(map[int]int)
	marbleList := []int{0}

	placeIndex := 1
	player := -1

	for i := 1; i <= nosMarbles; i++ {
		player = (player + 1) % nosPlayers

		if i%23 == 0 {
			playerScoreMap[player] += i

			// currentIndex := (placeIndex - 7 - 2) % len(marbleList)
			currentIndex := (placeIndex - 7 - 2)
			if currentIndex < 0 {
				currentIndex = len(marbleList) + currentIndex
			}
			// fmt.Println("placeIndex =", placeIndex)
			playerScoreMap[player] += marbleList[currentIndex]
			marbleList = RemoveAtIndex(marbleList, currentIndex)
			placeIndex = currentIndex
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
	return playerScoreMap
}
