package advent

import (
	"fmt"
)

type node struct {
	nosChildren int
	nosMetadata int
	mdEntries   []int
	children    []node
}

// MemoryManeuver is a func
func MemoryManeuver() {
	fmt.Println("---Memory Maneuver---")
}
