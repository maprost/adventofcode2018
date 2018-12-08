package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/maprost/adventofcode2018/golib"
)

type Point struct {
	id int
	x  int
	y  int
}

func main() {
	numbers := golib.Read("day06/task02/input_35237.txt")

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

	fmt.Println("Save areas: ", saveAreaCount)
}
