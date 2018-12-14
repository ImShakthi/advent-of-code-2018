package advent

import (
	"fmt"
	"strconv"
)

// CalcFrequency is to calc frequency
func CalcFrequency() {
	fmt.Println("Frequency calculator")

	data := ReadDataFromFile("day_one.txt")
	noL := len(data)

	freqMap := make(map[int]int)
	sum := 0

	arr := []int{}
	for index := 0; index < len(data); index++ {
		num, _ := strconv.Atoi(data[index])
		arr = append(arr, num)
	}
	//fmt.Println(arr)
	for i := 0; i < noL; i++ {
		sum = sum + arr[i]
		freqMap[sum]++
	}
	fmt.Println("sum =", sum)

	for i := 0; ; {
		sum = sum + arr[i]
		if freqMap[sum] == 1 {
			fmt.Println(">>>> Freq =", sum)
			break
		}
		i++
		i = i % noL
		freqMap[sum]++
	}
}
