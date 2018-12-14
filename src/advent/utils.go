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

	if file, err := os.Open("/Users/sakthivel/go/src/advent/inputs/" + fileName); err != nil {
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
