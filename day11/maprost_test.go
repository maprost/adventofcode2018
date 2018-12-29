package day11

import (
	"fmt"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
	"github.com/maprost/testbox/must"
)

const (
	size    = 300
	input01 = "input01_19,41.txt"
	input02 = "input02_237,284,11.txt"
)

func TestTask01(t *testing.T) {
	golib.ShowDebug = false
	serialIn, result := golib.Reads(input01)
	serial := golib.ToInt(serialIn[0])
	cellSize := 3

	powerGrid := calculatePowerGrid(size, size, serial)
	x, y, lvl := highestPowerLevel(powerGrid, cellSize)
	fmt.Println("Result: ", x, ",", y)
	must.BeEqual(t, fmt.Sprintf("%d,%d", x, y), result)

	must.BeEqual(t, powerLevel(powerGrid, x, y, cellSize), lvl)
	showPowerGrid(powerGrid, x, y, cellSize)
}

func TestTask02(t *testing.T) {
	golib.ShowDebug = false
	serialIn, result := golib.Reads(input02)
	serial := golib.ToInt(serialIn[0])

	powerGrid := calculatePowerGrid(size, size, serial)
	x, y, lvl, cellSize := highestCellPowerLevel(powerGrid)
	must.BeEqual(t, fmt.Sprintf("%d,%d,%d", x, y, cellSize), result)

	must.BeEqual(t, powerLevel(powerGrid, x, y, cellSize), lvl)
	showPowerGrid(powerGrid, x, y, cellSize)
}

// ================== tests =============================

func TestCalculatePower(t *testing.T) {
	golib.ShowDebug = true

	must.BeEqual(t, calcPower(122, 79, 57), -5)
	must.BeEqual(t, calcPower(217, 196, 39), 0)
	must.BeEqual(t, calcPower(101, 153, 71), 4)
}

func TestSerial18CellSize3(t *testing.T) {
	golib.ShowDebug = false
	x := 33
	y := 45
	serial := 18
	lvl := 29
	cellSize := 3

	must.BeEqual(t, calcPower(x, y, serial), 4)
	must.BeEqual(t, calcPower(x+1, y, serial), 4)
	must.BeEqual(t, calcPower(x+2, y, serial), 4)

	must.BeEqual(t, calcPower(x, y+1, serial), 3)
	must.BeEqual(t, calcPower(x+1, y+1, serial), 3)
	must.BeEqual(t, calcPower(x+2, y+1, serial), 4)

	must.BeEqual(t, calcPower(x, y+2, serial), 1)
	must.BeEqual(t, calcPower(x+1, y+2, serial), 2)
	must.BeEqual(t, calcPower(x+2, y+2, serial), 4)

	powerGrid := calculatePowerGrid(size, size, serial)
	showPowerGrid(powerGrid, x, y, cellSize)

	must.BeEqual(t, powerLevel(powerGrid, x, y, cellSize), lvl)

	fX, fY, fLvl := highestPowerLevel(powerGrid, cellSize)
	must.BeEqual(t, fX, x)
	must.BeEqual(t, fY, y)
	must.BeEqual(t, fLvl, lvl)
}

func TestSerial42CellSize3(t *testing.T) {
	golib.ShowDebug = false
	x := 21
	y := 61
	serial := 42
	lvl := 30
	cellSize := 3

	must.BeEqual(t, calcPower(x, y, serial), 4)
	must.BeEqual(t, calcPower(x+1, y, serial), 3)
	must.BeEqual(t, calcPower(x+2, y, serial), 3)

	must.BeEqual(t, calcPower(x, y+1, serial), 3)
	must.BeEqual(t, calcPower(x+1, y+1, serial), 3)
	must.BeEqual(t, calcPower(x+2, y+1, serial), 4)

	must.BeEqual(t, calcPower(x, y+2, serial), 3)
	must.BeEqual(t, calcPower(x+1, y+2, serial), 3)
	must.BeEqual(t, calcPower(x+2, y+2, serial), 4)

	powerGrid := calculatePowerGrid(size, size, serial)
	showPowerGrid(powerGrid, x, y, cellSize)

	must.BeEqual(t, powerLevel(powerGrid, x, y, cellSize), lvl)

	fX, fY, fLvl := highestPowerLevel(powerGrid, cellSize)
	must.BeEqual(t, fX, x)
	must.BeEqual(t, fY, y)
	must.BeEqual(t, fLvl, lvl)
}

func TestSerial42CellSize12(t *testing.T) {
	golib.ShowDebug = false
	x := 90
	y := 269
	serial := 42
	lvl := 119
	cellSize := 12

	powerGrid := calculatePowerGrid(size, size, serial)
	showPowerGrid(powerGrid, x, y, cellSize)

	must.BeEqual(t, powerLevel(powerGrid, x, y, cellSize), lvl)

	fX, fY, fLvl := highestPowerLevel(powerGrid, cellSize)
	must.BeEqual(t, fX, x)
	must.BeEqual(t, fY, y)
	must.BeEqual(t, fLvl, lvl)
}

// ================== program =============================

func calculatePowerGrid(maxX int, maxY int, serialNumber int) [][]int {
	powerGrid := make([][]int, maxY)

	for y := 0; y < maxY; y++ {
		powerGrid[y] = make([]int, maxX)

		for x := 0; x < maxX; x++ {
			powerGrid[y][x] = calcPower(x+1, y+1, serialNumber)
		}
	}

	return powerGrid
}

func highestCellPowerLevel(powerGrid [][]int) (x, y, power, cellSize int) {
	highestPowerLvl := 0
	foundX := 0
	foundY := 0
	foundCellSize := 0

	for c := 1; c < size; c++ {
		x, y, powerLvl := highestPowerLevel(powerGrid, c)

		if powerLvl > highestPowerLvl {
			highestPowerLvl = powerLvl
			foundX = x
			foundY = y
			foundCellSize = c
			fmt.Println("x:", x, " y:", y, " power:", powerLvl, " cellSize:", c)
		}
	}

	return foundX, foundY, highestPowerLvl, foundCellSize
}

func highestPowerLevel(powerGrid [][]int, cellSize int) (x, y, power int) {
	highestPowerLvl := 0
	foundX := 0
	foundY := 0

	for y := 0; y < len(powerGrid)-cellSize; y++ {
		for x := 0; x < len(powerGrid[y])-cellSize; x++ {
			powerLvl := powerLevel(powerGrid, x+1, y+1, cellSize)

			if powerLvl > highestPowerLvl {
				highestPowerLvl = powerLvl
				foundX = x + 1
				foundY = y + 1
			}
		}
	}

	return foundX, foundY, highestPowerLvl
}

func showPowerGrid(powerGrid [][]int, rX int, rY int, cellSize int) [][]int {
	golib.ShowDebug = false
	rY -= 2
	rX -= 2
	size := cellSize + 2

	for y := rY; y < len(powerGrid) && y < rY+size; y++ {
		for x := rX; x < len(powerGrid[y]) && x < rX+size; x++ {
			fmt.Printf("%2d ", powerGrid[y][x])
		}
		fmt.Println()
	}

	return powerGrid
}

func calcPower(x int, y int, serialNumber int) int {
	rackID := x + 10
	powerLevel := rackID * y
	plusSerial := powerLevel + serialNumber
	multiRackID := plusSerial * rackID

	hundredsDigit := 0
	// the third digit from right
	if multiRackID >= 100 {
		hundredsDigit = int((multiRackID % 1000) / 100)
	}
	power := hundredsDigit - 5

	golib.Debugf("(((%d + 10) * %d) + %d) = %d getThirdDigit ==> %d - 5 = %2d\n", x, y, serialNumber, multiRackID, hundredsDigit, power)
	return power
}

func powerLevel(powerGrid [][]int, rX int, rY int, cellSize int) int {
	rY -= 1
	rX -= 1
	size := cellSize
	power := 0

	for y := rY; y < len(powerGrid) && y < rY+size; y++ {
		for x := rX; x < len(powerGrid[y]) && x < rX+size; x++ {
			power += powerGrid[y][x]
		}
	}

	return power
}
