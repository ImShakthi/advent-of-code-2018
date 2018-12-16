package advent

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	return max(-a, -b)
}

func calculateDistance(p1, p2 coord) int {
	return abs(p1.x-p2.x) + abs(p1.y-p2.y)
}

func makeCoord(d string) coord {
	c := strings.Split(d, ", ")
	x, _ := strconv.Atoi(c[1])
	y, _ := strconv.Atoi(c[0])
	return coord{x: x, y: y}
}

// ChronalCoordinates is a function
func ChronalCoordinates() {
	fmt.Println("---Chronal Coordinates---")

	data := ReadDataFromFile("day_six.txt")

	// Find the outer border:
	var outerTop, outerBottom, outerLeft, outerRight int
	outerBottom = math.MaxInt32
	outerLeft = math.MaxInt32
	P := []coord{}
	for _, d := range data {
		c := makeCoord(d)
		P = append(P, c)
		outerTop = max(outerTop, c.y)
		outerBottom = min(outerBottom, c.y)
		outerLeft = min(outerLeft, c.x)
		outerRight = max(outerRight, c.x)
	}

	fmt.Println("outerBottom =", outerBottom, ", outerTop =", outerTop, ", outerLeft =", outerLeft, " outerRight =", outerRight)
	// Map coordinates to closest points and count:
	countP := map[coord]int{}
	canvas := map[coord]int{}
	for y := outerBottom; y <= outerTop; y++ {
		for x := outerLeft; x <= outerRight; x++ {
			c := coord{x: x, y: y}
			canvas[c] = 0
			minDist := math.MaxInt32
			var minP coord
			for _, p := range P {
				dist := calculateDistance(c, p)
				canvas[c] += dist
				if dist == minDist {
					minP = coord{x: math.MinInt32, y: math.MinInt32}
				} else if dist < minDist {
					minP = p
					minDist = dist
				}
			}

			if minP.x != math.MinInt32 && minP.y != math.MinInt32 {
				if countP[minP] != math.MinInt32 {
					countP[minP] = countP[minP] + 1
				}
				if y == outerTop || y == outerBottom || x == outerLeft || x == outerRight {
					countP[minP] = math.MinInt32
				}
			}
		}
	}

	// Part 1:
	fmt.Println("Part 1:")
	largestArea := math.MinInt32
	var largestCoord coord
	for c, area := range countP {
		if area > largestArea {
			largestArea = area
			largestCoord = c
		}
	}
	fmt.Printf("Point with largest area: %v\n", largestCoord)
	fmt.Printf("The area is: %v\n", largestArea)

	// Part 2:
	fmt.Println("\nPart 2:")
	var region []coord
	for p, d := range canvas {
		if d < 10000 {
			region = append(region, p)
		}
	}
	fmt.Printf("Size of region: %v\n", len(region))

}
