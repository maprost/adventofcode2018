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
	input := golib.Read("day03/task01/input_112418.txt")

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

	// result
	fmt.Println("Overlaps: ", countOverlap(fabric))
}

func fill(fabric [][]int, id int, x int, y int) {
	if fabric[x][y] != empty {
		fabric[x][y] = overlap
	} else {
		fabric[x][y] = id
	}
}

func countOverlap(fabric [][]int) int {
	overlapCounter := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if fabric[i][j] == overlap {
				overlapCounter++
			}
		}
	}
	return overlapCounter
}
