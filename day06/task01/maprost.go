package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/maprost/adventofcode2018/golib"
)

const (
	multiplyPointsSameDistance = -1
)

type Point struct {
	id int
	x  int
	y  int
}

func main() {
	numbers := golib.Read("day06/task01/input_4475.txt")

	points := make([]Point, 0, len(numbers))
	maxX := 0
	maxY := 0
	for i, n := range numbers {
		ns := strings.Split(n, ", ")
		x, err := strconv.Atoi(ns[1])
		if err != nil {
			panic(err)
		}
		if maxX < x {
			maxX = x
		}

		y, err := strconv.Atoi(ns[0])
		if err != nil {
			panic(err)
		}
		if maxY < y {
			maxY = y
		}

		points = append(points, Point{id: i, x: x, y: y})
	}

	maxX++
	maxY++

	// init coords
	coords := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		coords[i] = make([]int, maxY)
	}

	// print coords
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			id := "."
			for _, p := range points {
				if x == p.x && y == p.y {
					id = strconv.Itoa(p.id)
				}
			}
			fmt.Printf("%s ", id)
		}
		fmt.Println()
	}

	// calc Manhattan distance
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {

			minDistance := math.MaxInt32
			pId := multiplyPointsSameDistance

			for _, p := range points {
				dist := int(math.Abs(float64(x-p.x)) + math.Abs(float64(y-p.y)))

				if dist < minDistance {
					minDistance = dist
					pId = p.id

				} else if dist == minDistance {
					pId = multiplyPointsSameDistance
				}
			}

			coords[x][y] = pId
			fmt.Printf("%02d ", pId)
		}
		fmt.Println()
	}

	// count all id's that are on the edge
	infinitIds := make(map[int]struct{})
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			if x == 0 || y == 0 || x == maxX-1 || y == maxY-1 {
				infinitIds[coords[x][y]] = struct{}{}
			}
		}
	}

	// calculate for all ids the area they have
	areaSizes := make(map[int]int)
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {
			id := coords[x][y]
			if _, ok := infinitIds[id]; !ok {
				dist := areaSizes[id]
				areaSizes[id] = dist + 1
			}
		}
	}

	// search largest area
	largestSize := 0
	for id, size := range areaSizes {
		fmt.Printf("id: %d -> %d\n", id, size)
		if size > largestSize {
			largestSize = size
		}
	}

	fmt.Printf("largest Size: %d\n", largestSize)
}
