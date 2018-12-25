package advent

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// ReadDataFromFile is to read data from file
func ReadDataFromFile(fileName string) (data []string) {
	fmt.Println("Reading data from file '" + fileName + "'...")

	if file, err := os.Open("/Users/sakthivel/go/advent-of-code-2018/src/advent/inputs/" + fileName); err != nil {
		log.Fatal(err)
		defer file.Close()
	} else if scanner := bufio.NewScanner(file); scanner == nil {
		log.Fatal(" scanner is nil")
	} else {
		data = []string{}
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
		return data
	}
	return nil
}

// ConvInt method to convert string into int
func ConvInt(data string) int {
	number, _ := strconv.Atoi(data)
	return number
}

// RemoveAtIndex to remove element in the given list
func RemoveAtIndex(list []int, index int) []int {
	target := []int{}
	target = append(target, (list[:index])...)
	if (index + 1) <= len(list) {
		target = append(target, (list[index+1:])...)
	}
	return target
}

// AppendAtIndex to append element in the given list
func AppendAtIndex(list []int, index int, value int) []int {

	target := []int{}
	target = append(target, (list[:index])...)
	target = append(target, value)

	if (index + 1) <= len(list) {
		target = append(target, (list[index:])...)
	}
	return target
}
