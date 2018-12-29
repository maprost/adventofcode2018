package day06

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
)

const (
	multiplyPointsSameDistance = -1
	input01                    = "input01_4475.txt"
	input02                    = "input02_35237.txt"
)

type Point struct {
	id int
	x  int
	y  int
}

func TestTask01(t *testing.T) {
	numbers, result := golib.Reads(input01)

	points, maxX, maxY := calcPoints(numbers)

	// init coords
	coords := make([][]int, maxX)
	for i := 0; i < maxX; i++ {
		coords[i] = make([]int, maxY)
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

	golib.Equal(t, "largest Size: ", largestSize, result)
}

func TestTask02(t *testing.T) {
	numbers, result := golib.Reads(input02)

	points, maxX, maxY := calcPoints(numbers)

	// calc save areas
	saveAreaCount := 0
	for x := 0; x < maxX; x++ {
		for y := 0; y < maxY; y++ {

			sum := 0
			for _, p := range points {
				dist := int(math.Abs(float64(x-p.x)) + math.Abs(float64(y-p.y)))
				sum += dist
			}

			if sum < 10000 {
				saveAreaCount++
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}

	golib.Equal(t, "Save areas: ", saveAreaCount, result)
}

func calcPoints(numbers []string) ([]Point, int, int) {
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

	fmt.Println()

	return points, maxX, maxY
}
