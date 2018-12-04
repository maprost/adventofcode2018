package main

import (
	"fmt"

	"github.com/maprost/adventofcode2018/golib"
)

const (
	size    = 1000
	empty   = 0
	overlap = -1
)

func main() {
	input := golib.Read("day03/task02/input_560.txt")

	// init fabric
	fabric := make([][]int, size)
	for i := 0; i < size; i++ {
		fabric[i] = make([]int, size)
	}

	// fill fabric
	for _, in := range input {
		id, leftEdge, topEdge, wide, height := golib.SplitFabricInput(in)

		for w := leftEdge; w < leftEdge+wide; w++ {
			for h := topEdge; h < topEdge+height; h++ {
				fill(fabric, id, w, h)
			}
		}
	}

	// check overlaps
	for _, in := range input {
		id, leftEdge, topEdge, wide, height := golib.SplitFabricInput(in)
		noOverlaps := true

		for w := leftEdge; w < leftEdge+wide; w++ {
			for h := topEdge; h < topEdge+height; h++ {
				noOverlaps = noOverlaps && fabric[w][h] == id
			}
		}

		if noOverlaps {
			fmt.Println("No Overlaps: ", id)
		}
	}
}

func fill(fabric [][]int, id int, x int, y int) {
	if fabric[x][y] != empty {
		fabric[x][y] = overlap
	} else {
		fabric[x][y] = id
	}
}
