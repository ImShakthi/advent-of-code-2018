package advent

import (
	"fmt"
	"regexp"
	"sort"
)

//NightWatch struct
type NightWatch struct {
	guardID        int
	totalTimeSlept int
	times          []timeStamp
}

type timeStamp struct {
	sleptTime int
	wokeTime  int
}

// [1518-11-23 00:04] Guard #1663 begins shift
// [1518-11-23 00:56] wakes up
// [1518-11-23 00:19] falls asleep
// (?P<>[0-9]*)

const timeStampPattern = "\\[[0-9]*-[0-9]*-[0-9]*\\s[0-9]*:(?P<MINUTES>[0-9]*)\\]\\s"
const guardPattern = "Guard #(?P<GUARDID>[0-9]*) begins shift"
const wakesup = "wakes up"
const fallAsleep = "falls asleep"

var tsMetaData = regexp.MustCompile(timeStampPattern + guardPattern)
var wsMetaData = regexp.MustCompile(timeStampPattern + wakesup)
var fsMetaData = regexp.MustCompile(timeStampPattern + fallAsleep)

// ReposeRecord is function
func ReposeRecord() {
	fmt.Println("ReposeRecord")
	data := ReadDataFromFile("day_four.txt")
	noL := len(data)
	sort.Strings(data)

	guardMap := make(map[int]NightWatch)

	slept := 0
	guardID := 0

	currentWatch := NightWatch{}
	currentTS := timeStamp{}

	for index := 0; index < noL; index++ {
		//fmt.Println(data[index])
		input := data[index]
		if tsMetaData.Match([]byte(input)) {
			matches := tsMetaData.FindStringSubmatch(input)
			// minutes := matches[1]
			guardID = ConvInt(matches[2])
			currentWatch = guardMap[guardID]
			currentWatch.guardID = guardID
			currentTS = timeStamp{}
			//fmt.Println("guardID =", guardID)
		} else if fsMetaData.Match([]byte(input)) {
			matches := fsMetaData.FindStringSubmatch(input)
			slept = ConvInt(matches[1])
			currentTS.sleptTime = slept
			//fmt.Println("FALL minutes =", slept)
		} else if wsMetaData.Match([]byte(input)) {
			matches := wsMetaData.FindStringSubmatch(input)
			wokeup := ConvInt(matches[1])

			currentTS.wokeTime = wokeup
			currentWatch.totalTimeSlept += (currentTS.wokeTime - currentTS.sleptTime)
			currentWatch.times = append(currentWatch.times, currentTS)

			guardMap[guardID] = currentWatch
			//fmt.Println("WAKES minutes =", wokeup)
		}
	}
	//fmt.Println(guardMap)

	max := -1
	bestFitWatch := NightWatch{}
	for _, nightWatch := range guardMap {
		//fmt.Println("guardID =", guardID, ", Night watch=", nightWatch)
		if max < nightWatch.totalTimeSlept {
			bestFitWatch = nightWatch
			max = bestFitWatch.totalTimeSlept
		}
	}

	//mt.Println("Best watch to get in it is ", bestFitWatch)
	_, highFreqTime := getHighFreqTime(bestFitWatch)
	fmt.Println("highTimeFreq * guardID =", highFreqTime*bestFitWatch.guardID)

	fmt.Println("--------- PART TWO ----------")
	//fmt.Println("no of guard =", len(guardMap))

	maxFreqOfGuards := make(map[int][2]int)
	var freqHighFreqArr [2]int
	for _, nightWatch := range guardMap {
		frequency, highFreqTime := getHighFreqTime(nightWatch)
		freqHighFreqArr[0] = frequency
		freqHighFreqArr[1] = highFreqTime
		maxFreqOfGuards[nightWatch.guardID] = freqHighFreqArr
	}
	//fmt.Println("maxFreqOfGuards =", maxFreqOfGuards)
	hfGuardID := -1
	hfTime := -1
	max = 0
	for guardID, freqHighFreqArr := range maxFreqOfGuards {
		if max < freqHighFreqArr[0] {
			max = freqHighFreqArr[0]
			hfTime = freqHighFreqArr[1]
			hfGuardID = guardID
		}
	}
	fmt.Println("Max Freq on same time =", hfGuardID*hfTime, ", hfGuardID =", hfGuardID, ",  hfTime=", hfTime)
}

func getHighFreqTime(bestFitWatch NightWatch) (int, int) {
	var frequencyArr [60]int
	for index := 0; index < len(bestFitWatch.times); index++ {
		ts := bestFitWatch.times[index]
		for i := ts.sleptTime; i < ts.wokeTime; i++ {
			frequencyArr[i]++
		}
	}
	//fmt.Println(counterArr)
	frequency := -1
	highFreqTime := -1
	for index := 0; index < len(frequencyArr); index++ {
		if frequency < frequencyArr[index] {
			frequency = frequencyArr[index]
			highFreqTime = index
		}
	}
	return frequency, highFreqTime
}
