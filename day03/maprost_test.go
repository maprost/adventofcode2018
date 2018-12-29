package day03

import (
	"strconv"
	"strings"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
)

const (
	size    = 1000
	empty   = 0
	overlap = -1

	input01 = "input01_112418.txt"
	input02 = "input02_560.txt"
)

func TestTask01(t *testing.T) {
	input, result := golib.Reads(input01)
	fabric := buildFabric(input)

	// result
	overlaps := countOverlap(fabric)
	golib.Equal(t, "Overlaps: ", overlaps, result)
}

func TestTask02(t *testing.T) {
	input, result := golib.Reads(input02)
	fabric := buildFabric(input)

	// check overlaps
	for _, in := range input {
		id, leftEdge, topEdge, wide, height := splitInput(in)
		noOverlaps := true

		for w := leftEdge; w < leftEdge+wide; w++ {
			for h := topEdge; h < topEdge+height; h++ {
				noOverlaps = noOverlaps && fabric[w][h] == id
			}
		}

		if noOverlaps {
			golib.Equal(t, "No Overlaps: ", id, result)
		}
	}
}

func buildFabric(input []string) [][]int {
	// init fabric
	fabric := make([][]int, size)
	for i := 0; i < size; i++ {
		fabric[i] = make([]int, size)
	}

	// fill fabric
	for _, in := range input {
		id, leftEdge, topEdge, wide, height := splitInput(in)

		for w := leftEdge; w < leftEdge+wide; w++ {
			for h := topEdge; h < topEdge+height; h++ {
				fill(fabric, id, w, h)
			}
		}
	}

	return fabric
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

func splitInput(in string) (id, leftEdge, topEdge, wide, height int) {
	var err error
	inArr := strings.Split(in, " ")

	id, err = strconv.Atoi(strings.TrimLeft(inArr[0], "#"))
	if err != nil {
		panic(err)
	}

	edgePoints := strings.Split(inArr[2], ",")
	leftEdge, err = strconv.Atoi(edgePoints[0])
	if err != nil {
		panic(err)
	}

	topEdge, err = strconv.Atoi(strings.TrimRight(edgePoints[1], ":"))
	if err != nil {
		panic(err)
	}

	size := strings.Split(inArr[3], "x")
	wide, err = strconv.Atoi(size[0])
	if err != nil {
		panic(err)
	}

	height, err = strconv.Atoi(size[1])
	if err != nil {
		panic(err)
	}

	return
}
