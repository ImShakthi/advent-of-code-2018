package advent

import (
	"fmt"
	"regexp"
	"sort"
)

var graphDataPattern = "Step\\s(?P<P1>[A-Z])\\smust be finished before step\\s(?P<P2>[A-Z])\\scan begin."
var metaData = regexp.MustCompile(graphDataPattern)

var graphPath = ""
var lastEdge = ""
var counter = 0

//TheSumofItsParts is func
func TheSumofItsParts() {
	fmt.Println("---The Sum of Its Parts---")
	data := ReadDataFromFile("day_seven.txt")
	noL := len(data)

	isFirst := true
	firstV := " "
	graph := make(map[string][]string)
	for index := 0; index < noL; index++ {
		claim := parseGraphInput(data[index])
		//fmt.Println(claim)
		edges := graph[claim[0]]
		edges = append(edges, claim[1])
		graph[claim[0]] = edges
		if isFirst {
			isFirst = false
			firstV = claim[0]
		}
	}
	// DGAFIKRTZJLPMUQXSWR
	// DLPZKRTFJIMAUSWGQXS
	// DGAFIKRTZJLPMUQXSWR
	for key := range graph {
		edges := graph[key]
		sort.Strings(edges)
		graph[key] = edges
	}
	//fmt.Println(graph)
	var visited = make(map[string]bool)

	printGraph(graph, firstV, visited)
	graphPath += lastEdge
	fmt.Println(graphPath, ", dataLen =", len(graphPath), ", maplen =", len(graph))

}

func printGraph(graph map[string][]string, currentV string, visited map[string]bool) {
	if !visited[currentV] {
		graphPath += currentV
		visited[currentV] = true
	}
	edges := graph[currentV]
	fmt.Println("counter =", counter, "key =", currentV, ", edges =", edges)
	counter++
	for index := range edges {
		value := edges[index]
		fmt.Println(">> value =", value)
		if !visited[value] {
			visited[value] = true
			graphPath += value
			if len(graph[value]) != 0 {
				printGraph(graph, value, visited)
			}
		}
	}
}

// DLPZKRTFJIMAUSWGQXS
// DLPZKRTFJIMAUSWGQXS
func parseGraphInput(inputStr string) []string {
	pts := []string{}
	matches := metaData.FindStringSubmatch(inputStr)
	pts = append(pts, matches[1])
	pts = append(pts, matches[2])

	return pts
}
