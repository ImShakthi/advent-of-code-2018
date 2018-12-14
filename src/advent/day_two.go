package advent

import (
	"fmt"
)

// CalcThreshold is to calc Threshold
func CalcThreshold() {
	fmt.Println("Threshold calculator....")

	prodIDArr := ReadDataFromFile("day_two.txt")
	noL := len(prodIDArr)

	// Part One
	freqOfTwo := 0
	freqOfThree := 0
	for index := 0; index < len(prodIDArr); index++ {
		charMap := getCharFreq(prodIDArr[index])
		// check for twice
		if occurs(charMap, 2) {
			freqOfTwo++
		}
		// check for thrice
		if occurs(charMap, 3) {
			freqOfThree++
		}
	}
	fmt.Println("freq of two =", freqOfTwo, ", and three =", freqOfThree, ", total threshold frequency = ", freqOfThree*freqOfTwo, ", len =", noL)

	// Part Two
	for index := 0; index < len(prodIDArr); index++ {
		currentProdID := prodIDArr[index]
		for j := index; j < len(prodIDArr); j++ {
			mismatchMoreThanOnce := false
			mismatchMoreOnce := false
			secProdID := prodIDArr[j]

			breakIndex := 0
			for k := 0; k < len(secProdID); k++ {
				if currentProdID[k] != secProdID[k] {
					if mismatchMoreOnce {
						mismatchMoreThanOnce = true
						break
					}
					breakIndex = k
					mismatchMoreOnce = true
				}
			}
			if mismatchMoreOnce && !mismatchMoreThanOnce {
				result := secProdID[0:breakIndex] + secProdID[breakIndex+1:len(secProdID)]
				fmt.Println("Part TWO >>> Found ", secProdID, ", result =", result)

			}
		}
	}

}

func occurs(charMap map[byte]int, occurence int) bool {
	for _, value := range charMap {
		if value == occurence {
			return true
		}
	}
	return false
}

func getCharFreq(prodID string) map[byte]int {
	charMap := make(map[byte]int)

	for index := 0; index < len(prodID); index++ {
		key := prodID[index]
		charMap[key]++
	}
	return charMap
}
