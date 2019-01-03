package day13

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/maprost/adventofcode2018/golib"
	"github.com/maprost/testbox/must"
	"github.com/maprost/testbox/should"
)

const (
	UP    = "^"
	DOWN  = "v"
	LEFT  = "<"
	RIGHT = ">"

	I_LEFT     = 0
	I_STRAIGHT = 1
	I_RIGHT    = 2

	input01 = "input01_.txt"
	input02 = "input02_.txt"
)

func TestTask01(t *testing.T) {
	golib.ShowDebug = false

	input := golib.Read(input01)
	m := loadMap(input)

	fmt.Println(m)
	fmt.Println("Coords: ", findCollision(m))
}

func TestTask02(t *testing.T) {

}

// ================== tests =============================

func TestExample1(t *testing.T) {
	golib.ShowDebug = true

	m := loadMap([]string{
		"|",
		"v",
		"|",
		"|",
		"|",
		"^",
		"|",
	})

	fmt.Println(m)
	must.BeEqual(t, findCollision(m), "0,3")
}

func TestExample2(t *testing.T) {
	golib.ShowDebug = true

	m := loadMap([]string{
		"|",
		"v",
		"|",
		"|",
		"^",
		"|",
	})

	fmt.Println(m)
	must.BeEqual(t, findCollision(m), "0,3")
}

func TestExample3(t *testing.T) {
	golib.ShowDebug = true

	m := loadsMap(`       
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/
`)

	fmt.Println(m)
	must.BeEqual(t, findCollision(m), "7,4")
}

func TestMove(t *testing.T) {
	golib.ShowDebug = false

	checkIntersectionMove := func(t testing.TB, in []string, intersection int, exp string) {
		t.Helper()
		m := loadMap(in)

		for _, arrow := range m.arrows {
			delete(m.arrows, arrowPos(arrow.x, arrow.y))

			arrow.intersection = intersection
			m.arrows[arrowPos(arrow.x, arrow.y)] = arrow
		}

		m.Move()
		should.BeEqual(t, m.String(), exp)
	}
	checkMove := func(t testing.TB, in []string, exp string) {
		t.Helper()
		checkIntersectionMove(t, in, 0, exp)
	}

	// right moves
	checkMove(t, []string{"->-"}, "-->\n")
	checkMove(t, []string{"->\\"}, "--v\n")
	checkMove(t, []string{"->/"}, "--^\n")
	checkIntersectionMove(t, []string{"->+"}, I_LEFT, "--^\n")
	checkIntersectionMove(t, []string{"->+"}, I_STRAIGHT, "-->\n")
	checkIntersectionMove(t, []string{"->+"}, I_RIGHT, "--v\n")

	// left moves
	checkMove(t, []string{"-<-"}, "<--\n")
	checkMove(t, []string{"\\<-"}, "^--\n")
	checkMove(t, []string{"/<-"}, "v--\n")
	checkIntersectionMove(t, []string{"+<-"}, I_LEFT, "v--\n")
	checkIntersectionMove(t, []string{"+<-"}, I_STRAIGHT, "<--\n")
	checkIntersectionMove(t, []string{"+<-"}, I_RIGHT, "^--\n")

	// up moves
	checkMove(t, []string{"|", "^"}, "^\n|\n")
	checkMove(t, []string{"/", "^"}, ">\n|\n")
	checkMove(t, []string{"\\", "^"}, "<\n|\n")
	checkIntersectionMove(t, []string{"+", "^"}, I_LEFT, "<\n|\n")
	checkIntersectionMove(t, []string{"+", "^"}, I_STRAIGHT, "^\n|\n")
	checkIntersectionMove(t, []string{"+", "^"}, I_RIGHT, ">\n|\n")

	// down moves
	checkMove(t, []string{"v", "|"}, "|\nv\n")
	checkMove(t, []string{"v", "/"}, "|\n<\n")
	checkMove(t, []string{"v", "\\"}, "|\n>\n")
	checkIntersectionMove(t, []string{"v", "+"}, I_LEFT, "|\n>\n")
	checkIntersectionMove(t, []string{"v", "+"}, I_STRAIGHT, "|\nv\n")
	checkIntersectionMove(t, []string{"v", "+"}, I_RIGHT, "|\n<\n")
}

func TestCollision(t *testing.T) {

}

// ================== program =============================

type Arrow struct {
	intersection int
	x            int
	y            int
	symbol       string
}

type Map struct {
	xMax       int
	yMax       int
	background [][]string
	arrows     map[string]Arrow
}

func (m Map) String() string {
	result := ""
	for y := 0; y < m.yMax; y++ {
		for x := 0; x < m.xMax; x++ {
			if arrow, ok := m.arrows[arrowPos(x, y)]; ok {
				result += arrow.symbol
			} else {
				result += m.background[y][x]
			}
		}
		result += "\n"
	}
	return result
}

func (m *Map) Move() (bool, string) {
	nextArrows := make(map[string]Arrow)
	alright := true
	brokenCoords := ""

	for _, arrow := range m.arrows {
		nextX := arrow.x
		nextY := arrow.y
		nextSymbol := arrow.symbol
		nextIntersection := arrow.intersection

		backslashCurveSymbol := ""
		slashCurveSymbol := ""
		iLeftSymbol := ""
		iRightSymbol := ""

		// ->-
		if arrow.symbol == RIGHT {
			nextX++
			backslashCurveSymbol = DOWN // -->\ => ---V
			slashCurveSymbol = UP       // -->/ => ---^
			iLeftSymbol = UP            // -->+  => ---^
			iRightSymbol = DOWN         // -->+  => ---V
		}

		// -<-
		if arrow.symbol == LEFT {
			nextX--
			backslashCurveSymbol = UP // \<-- => ^---
			slashCurveSymbol = DOWN   // /<-- => V---
			iLeftSymbol = DOWN        // +<-- => V---
			iRightSymbol = UP         // +<-- => ^---
		}

		// |
		// V
		// |
		if arrow.symbol == DOWN {
			nextY++
			backslashCurveSymbol = RIGHT // \
			slashCurveSymbol = LEFT      // /
			iLeftSymbol = RIGHT
			iRightSymbol = LEFT
		}

		// |
		// ^
		// |
		if arrow.symbol == UP {
			nextY--
			backslashCurveSymbol = LEFT // \
			slashCurveSymbol = RIGHT    // /
			iLeftSymbol = LEFT
			iRightSymbol = RIGHT
		}

		// curve:
		if m.background[nextY][nextX] == "\\" {
			nextSymbol = backslashCurveSymbol
		}
		if m.background[nextY][nextX] == "/" {
			nextSymbol = slashCurveSymbol
		}
		// crossing:
		if m.background[nextY][nextX] == "+" {
			switch arrow.intersection {
			case I_LEFT:
				nextSymbol = iLeftSymbol
			case I_RIGHT:
				nextSymbol = iRightSymbol
			}
			nextIntersection = (arrow.intersection + 1) % 3
		}

		coords := arrowPos(nextX, nextY)

		// check collision
		if _, ok := nextArrows[coords]; ok {
			alright = false
			brokenCoords += coords
		}
		//if _, ok := m.arrows[coords]; ok {
		//	alright = false
		//	brokenCoords += coords
		//}

		// add new arrow
		nextArrows[coords] = Arrow{intersection: nextIntersection, x: nextX, y: nextY, symbol: nextSymbol}
		delete(m.arrows, coords)
	}

	m.arrows = nextArrows
	return alright, brokenCoords
}

func loadsMap(input string) Map {
	return loadMap(strings.Split(input, "\n"))
}

func loadMap(input []string) Map {
	xMax := 0
	for _, s := range input {
		if len(s) > xMax {
			xMax = len(s)
		}
	}

	yMax := len(input)

	m := Map{
		xMax:       xMax,
		yMax:       yMax,
		background: make([][]string, yMax),
		arrows:     make(map[string]Arrow),
	}
	for y, line := range input {
		m.background[y] = make([]string, xMax)

		for x, r := range line {
			s := string(r)

			if s == LEFT || s == RIGHT {
				m.background[y][x] = "-"
				m.arrows[arrowPos(x, y)] = Arrow{intersection: 0, x: x, y: y, symbol: s}

			} else if s == UP || s == DOWN {
				m.background[y][x] = "|"
				m.arrows[arrowPos(x, y)] = Arrow{intersection: 0, x: x, y: y, symbol: s}

			} else {
				m.background[y][x] = s
			}
		}
	}

	return m
}

func arrowPos(x int, y int) string {
	return strconv.Itoa(x) + "," + strconv.Itoa(y)
}

func findCollision(m Map) string {
	ok := true
	coords := ""

	for ok {
		ok, coords = m.Move()
		golib.Debugln(m)
	}

	return coords
}
