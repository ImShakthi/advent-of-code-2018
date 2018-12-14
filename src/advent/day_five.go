package advent

import (
	"fmt"
	"strings"
	"unicode"
)

// AlchemicalReduction is function
func AlchemicalReduction() {
	fmt.Println("AlchemicalReduction")
	fileData := ReadDataFromFile("day_five.txt")
	//noL := len(fileData)
	data := fileData[0]

	//fmt.Println(" # of units after removing polymer =", getPolymerUnits(data))
	fmt.Println("mini polymer # =", getMiniPolymer(data))
}

func getPolymerUnits(data string) int {
	isPolymerPresent := true
	for isPolymerPresent {
		isPolymerPresent, data = removePolymer(data)
		//refmt.Println("After iteration >> ", isPolymerPresent, data)
	}
	return len(data)
}

func getMiniPolymer(data string) int {
	miniPolymer := 50001

	miniPolymerArr := []int{}

	for lowCaps, bigCaps := 'a', 'A'; lowCaps <= 'z'; {
		if strings.Contains(data, string(lowCaps)) || strings.Contains(data, string(bigCaps)) {
			fmt.Println("Started for ", string(lowCaps))
			updatedData := strings.Replace(data, string(lowCaps), "", -1)
			updatedData = strings.Replace(updatedData, string(bigCaps), "", -1)

			polymerUnit := getPolymerUnits(updatedData)
			miniPolymerArr = append(miniPolymerArr, polymerUnit)
			fmt.Println("Ended for ", string(lowCaps))
		} else {
			fmt.Println(">>> Skipped ", string(lowCaps))

		}
		lowCaps++
		bigCaps++
	}

	fmt.Println(miniPolymerArr)
	for index := 0; index < len(miniPolymerArr); index++ {
		if miniPolymerArr[index] < miniPolymer {
			miniPolymer = miniPolymerArr[index]
		}
	}
	return miniPolymer
}
func removePolymer(data string) (bool, string) {
	isPolymerPresent := false
	updatedData := ""

	var prevChar rune
	prevCharCaps := false
	allow := true
	for _, char := range data {
		currentChar := string(char)
		currentCharCaps := unicode.IsUpper(char)

		if currentCharCaps != prevCharCaps && unicode.ToLower(prevChar) == unicode.ToLower(char) && allow {
			isPolymerPresent = true
			allow = false
			//updatedData = updatedData[0:len(updatedData)]
			remove := string(updatedData[len(updatedData)-1])
			updatedData = strings.TrimSuffix(updatedData, remove)
		} else {
			updatedData += currentChar
			allow = true
		}

		//fmt.Println(index, data, updatedData)
		//fmt.Println(currentChar, currentCharCaps, prevChar, prevCharCaps)
		prevCharCaps = currentCharCaps
		prevChar = char
	}
	//dabCBAcaDA
	return isPolymerPresent, updatedData
}
